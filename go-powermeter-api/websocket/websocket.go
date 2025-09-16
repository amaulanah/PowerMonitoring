package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader untuk meng-upgrade koneksi HTTP ke WebSocket.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // Izinkan semua koneksi
}

// Fungsi Upgrade yang dipanggil dari main.go.
func Upgrade(w http.ResponseWriter, r *http.Request, pool *Pool) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Buat klien baru.
	client := &Client{
		Conn: conn,
		Pool: pool,
	}

	// Daftarkan klien baru ke pool.
	pool.Register <- client

	// Jalankan metode Read() dalam goroutine baru agar tidak memblokir.
	// Ini akan menjaga koneksi dan menangani disconnect.
	go client.Read()
}