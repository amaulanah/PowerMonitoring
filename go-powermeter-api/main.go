package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"

	// Pastikan path modul ini sesuai dengan file go.mod Anda
	"github.com/amaulanah/powermeterapi/auth"
	"github.com/amaulanah/powermeterapi/database"
	"github.com/amaulanah/powermeterapi/models"
	"github.com/amaulanah/powermeterapi/plc"
	"github.com/amaulanah/powermeterapi/websocket"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Proses 1: Polling cepat ke PLC dan broadcast ke WebSocket
func startPollingLoop(wsPool *websocket.Pool, dataForDB chan<- []models.PowerMeterReading, dbPool *pgxpool.Pool) {
	// Ticker sekarang berjalan setiap 200 milidetik
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 1. Ambil daftar perangkat dari database
			deviceIDs, err := database.GetAllDeviceIDs(dbPool)
			if err != nil {
				fmt.Println("Error saat mengambil daftar perangkat:", err)
				continue
			}

			if len(deviceIDs) == 0 {
				// Tidak perlu print pesan ini setiap 200ms
				continue
			}

			// 2. Baca data PLC berdasarkan daftar yang didapat
			readings, err := plc.ReadAllMetersData(deviceIDs)
			if err != nil {
				fmt.Println("Error saat membaca semua meter:", err)
				continue
			}

			if len(readings) > 0 {
				// 3. Langsung siarkan ke WebSocket untuk tampilan realtime
				wsPool.Broadcast <- readings

				// 4. Kirim data yang sama ke channel untuk diproses oleh database
				dataForDB <- readings
			}
		}
	}
}

// Proses 2: Penulis ke database yang berjalan lebih lambat
func startDatabaseWriter(dbPool *pgxpool.Pool, dataForDB <-chan []models.PowerMeterReading) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// Kita gunakan map untuk menyimpan hanya data terakhir dari setiap device
	latestReadingsMap := make(map[string]models.PowerMeterReading)

	for {
		select {
		case newReadings := <-dataForDB:
			// Saat data baru masuk, update map. Data lama akan tertimpa.
			for _, reading := range newReadings {
				latestReadingsMap[reading.DeviceID] = reading
			}

		case <-ticker.C:
			// Setiap 1 detik, proses map yang sudah terkumpul
			if len(latestReadingsMap) > 0 {
				// Ubah map kembali menjadi slice untuk dimasukkan ke database
				readingsToInsert := make([]models.PowerMeterReading, 0, len(latestReadingsMap))
				for _, reading := range latestReadingsMap {
					readingsToInsert = append(readingsToInsert, reading)
				}

				fmt.Printf("Menyimpan %d data pembacaan terakhir ke database...\n", len(readingsToInsert))
				err := database.InsertReadings(dbPool, readingsToInsert)
				if err != nil {
					fmt.Println("Error memasukkan ke DB:", err)
				}

				// Kosongkan map untuk interval 1 detik berikutnya
				latestReadingsMap = make(map[string]models.PowerMeterReading)
			}
		}
	}
}

func main() {
	// Muat variabel dari file .env di awal
	err := godotenv.Load()
	if err != nil {
		log.Println("Peringatan: Gagal memuat file .env")
	}

	dbPool, err := database.NewConnection()
	if err != nil {
		panic(err)
	}
	defer dbPool.Close()

	wsPool := websocket.NewPool()
	dataForDB := make(chan []models.PowerMeterReading, 10)

	// Jalankan TIGA proses di background
	go wsPool.Start()
	go startPollingLoop(wsPool, dataForDB, dbPool)
	go startDatabaseWriter(dbPool, dataForDB)

	router := gin.Default()

	// Konfigurasi CORS
	config := cors.DefaultConfig()
	// Ganti URL ini jika port frontend Anda berbeda
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	router.Use(cors.New(config))

	// --- Grup API yang diproteksi ---
	api := router.Group("/api")
	// Middleware untuk proteksi API akan kita tambahkan nanti jika perlu
	{
		// ENDPOINT BARU: Mengambil daftar semua perangkat
		api.GET("/devices", func(c *gin.Context) {
			deviceIDs, err := database.GetAllDeviceIDs(dbPool)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data perangkat"})
				return
			}
			c.JSON(http.StatusOK, deviceIDs)
		})

		// ENDPOINT BARU: Mengambil data historis untuk chart
		api.GET("/historical-data", func(c *gin.Context) {
			// Ambil parameter dari query URL
			deviceID := c.Query("deviceId")
			parameter := c.Query("parameter") // contoh: "ActivePowerTotal"
			interval := c.Query("interval")   // contoh: "hour"

			// Panggil fungsi dari database untuk mengambil data historis
			// (Fungsi ini perlu dibuat di dalam file database/database.go)
			data, err := database.GetHistoricalData(dbPool, deviceID, parameter, interval)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data historis"})
				return
			}
			c.JSON(http.StatusOK, data)
		})
	}

	router.POST("/login", func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
			return
		}

		user, err := database.GetUserByUsername(dbPool, req.Username)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau password salah"})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau password salah"})
			return
		}

		token, err := auth.GenerateJWT(user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	router.GET("/ws", func(c *gin.Context) {
		tokenStr := c.Query("token")
		if tokenStr == "" {
			return
		}
		_, err := auth.ValidateJWT(tokenStr)
		if err != nil {
			return
		}
		websocket.Upgrade(c.Writer, c.Request, wsPool)
	})

	fmt.Println("Server dimulai di http://localhost:1234")
	router.Run(":1234") // Menggunakan port 1234 sesuai kode Anda
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/jackc/pgx/v5/pgxpool"
// 	"github.com/joho/godotenv"
// 	"golang.org/x/crypto/bcrypt"

// 	// Ganti dengan path modul Anda
// 	"github.com/amaulanah/powermeterapi/auth"
// 	"github.com/amaulanah/powermeterapi/database"
// 	"github.com/amaulanah/powermeterapi/models"
// 	"github.com/amaulanah/powermeterapi/plc"
// 	"github.com/amaulanah/powermeterapi/websocket"
// )

// type LoginRequest struct {
// 	Username string `json:"username" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }

// func startPollingLoop(dbPool *pgxpool.Pool, wsPool *websocket.Pool) {
// 	// Timer yang berjalan setiap 1 detik
// 	ticker := time.NewTicker(1 * time.Second)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ticker.C:
// 			fmt.Println("Melakukan polling ke PLC...")
// 			// Lakukan pembacaan dari 20 meter
// 			// Untuk sekarang kita simulasikan
// 			reading, err := plc.ReadDataFromPlc("pm1")
// 			if err != nil {
// 				fmt.Println("Error membaca PLC:", err)
// 				continue
// 			}

// 			readings := []models.PowerMeterReading{*reading}

// 			// Masukkan ke DB
// 			err = database.InsertReadings(dbPool, readings)
// 			if err != nil {
// 				fmt.Println("Error memasukkan ke DB:", err)
// 			}

// 			// Siarkan ke WebSocket
// 			wsPool.Broadcast <- readings
// 		}
// 	}
// }

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Println("Peringatan: Gagal memuat file .env")
// 	}

// 	dbPool, err := database.NewConnection()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer dbPool.Close()

// 	wsPool := websocket.NewPool()
// 	go wsPool.Start()
// 	go startPollingLoop(dbPool, wsPool)

// 	router := gin.Default()

// 	// ENDPOINT BARU: POST /login
// 	router.POST("/login", func(c *gin.Context) {
// 		var req LoginRequest
// 		if err := c.ShouldBindJSON(&req); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
// 			return
// 		}

// 		// 1. Ambil data user dari DB
// 		user, err := database.GetUserByUsername(dbPool, req.Username)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau password salah"})
// 			return
// 		}

// 		// 2. Bandingkan password yang diberikan dengan hash di DB
// 		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau password salah"})
// 			return
// 		}

// 		// 3. Jika berhasil, buat token JWT
// 		token, err := auth.GenerateJWT(user.Username)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat token"})
// 			return
// 		}

// 		// 4. Kirim token ke client
// 		c.JSON(http.StatusOK, gin.H{"token": token})
// 	})

// 	// PROTEKSI ENDPOINT WEBSOCKET
// 	router.GET("/ws", func(c *gin.Context) {
// 		// Ambil token dari query parameter (contoh: ws://.../ws?token=xxxx)
// 		tokenStr := c.Query("token")
// 		if tokenStr == "" {
// 			return // Abaikan koneksi tanpa token
// 		}

// 		// Validasi token
// 		_, err := auth.ValidateJWT(tokenStr)
// 		if err != nil {
// 			// Token tidak valid, abaikan koneksi
// 			return
// 		}

// 		// Jika token valid, upgrade ke WebSocket
// 		websocket.Upgrade(c.Writer, c.Request, wsPool)
// 	})

// 	fmt.Println("Server dimulai di http://localhost:1234")
// 	router.Run(":1234")
// }
