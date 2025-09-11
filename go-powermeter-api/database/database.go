package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/amaulanah/powermeterapi/models" // GANTI DENGAN PATH MODUL ANDA
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Fungsi NewConnection tidak berubah
func NewConnection() (*pgxpool.Pool, error) {
	connString := "postgres://postgres:postgres@localhost:5432/powermeter_db" // GANTI KATA SANDI

	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Gagal membuat koneksi pool: %v\n", err)
		return nil, err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Gagal ping database: %v\n", err)
		return nil, err
	}

	fmt.Println("Berhasil terhubung ke database!")
	return pool, nil
}

// Fungsi InsertReadings diperbarui dengan semua kolom
func InsertReadings(pool *pgxpool.Pool, readings []models.PowerMeterReading) error {
	// Daftar kolom HARUS sama persis urutannya dengan di bawah
	columnNames := []string{
		"Timestamp", "DeviceId", "Active_Energy_Kwh", "Current_A", "Current_B", "Current_C", "Current_N",
		"Current_G", "Current_Avg", "Voltage_AB", "Voltage_BC", "Voltage_CA", "VoltageL_Avg", "Voltage_AN",
		"Voltage_BN", "Voltage_CN", "NA", "VoltageN_Avg", "Active_Power_A", "Active_Power_B", "Active_Power_C",
		"Active_Power_Total", "Reactive_Power_A", "Reactive_Power_B", "Reactive_Power_C", "Reactive_Power_Total",
		"Apparent_Power_A", "Apparent_Power_B", "Apparent_Power_C", "Apparent_Power_Total", "Power_Factor_A",
		"Power_Factor_B", "Power_Factor_C", "Power_Factor_Total", "Frequency",
	}

	rows := make([][]interface{}, len(readings))
	for i, r := range readings {
		// Urutan data HARUS sama persis dengan daftar kolom di atas
		rows[i] = []interface{}{
			r.Timestamp, r.DeviceID, r.Active_Energy_Kwh, r.Current_A, r.Current_B, r.Current_C, r.Current_N,
			r.Current_G, r.Current_Avg, r.Voltage_AB, r.Voltage_BC, r.Voltage_CA, r.VoltageL_Avg, r.Voltage_AN,
			r.Voltage_BN, r.Voltage_CN, r.NA, r.VoltageN_Avg, r.Active_Power_A, r.Active_Power_B, r.Active_Power_C,
			r.Active_Power_Total, r.Reactive_Power_A, r.Reactive_Power_B, r.Reactive_Power_C, r.Reactive_Power_Total,
			r.Apparent_Power_A, r.Apparent_Power_B, r.Apparent_Power_C, r.Apparent_Power_Total, r.Power_Factor_A,
			r.Power_Factor_B, r.Power_Factor_C, r.Power_Factor_Total, r.Frequency,
		}
	}

	_, err := pool.CopyFrom(
		context.Background(),
		pgx.Identifier{"readings"},
		columnNames,
		pgx.CopyFromRows(rows),
	)

	return err
}

// Struct untuk model User
type User struct {
	Username     string
	PasswordHash string
}

// Fungsi untuk mendapatkan data user berdasarkan username
func GetUserByUsername(pool *pgxpool.Pool, username string) (*User, error) {
	user := &User{}
	// Menggunakan parameterized query ($1) untuk mencegah SQL Injection
	err := pool.QueryRow(context.Background(), `SELECT "username", "password_hash" FROM users WHERE "username" = $1`, username).Scan(&user.Username, &user.PasswordHash)
	if err != nil {
		// Jika tidak ada baris (user tidak ditemukan), pgx akan mengembalikan pgx.ErrNoRows
		return nil, err
	}
	return user, nil
}

func GetAllDeviceIDs(pool *pgxpool.Pool) ([]string, error) {
	rows, err := pool.Query(context.Background(), `SELECT "Id" FROM devices ORDER BY "Id"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var deviceIDs []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		deviceIDs = append(deviceIDs, id)
	}
	return deviceIDs, nil
}

type HistoricalPoint struct {
	TimeBucket time.Time `json:"time"`
	AvgValue   float64   `json:"value"`
}

// Fungsi untuk mengambil data historis dengan agregasi
func GetHistoricalData(pool *pgxpool.Pool, deviceID string, parameter string, interval string) ([]HistoricalPoint, error) {
	// Validasi input untuk mencegah SQL Injection pada nama kolom dan interval
	// Ini adalah daftar kolom yang diizinkan untuk di-query
	allowedParameters := map[string]bool{
		"EnergyKWh": true, "ActivePowerTotal": true, "Frequency": true, // Tambahkan semua 29 parameter di sini
	}
	if !allowedParameters[parameter] {
		return nil, fmt.Errorf("parameter tidak valid: %s", parameter)
	}

	// Tentukan rentang waktu (contoh: 24 jam terakhir)
	startTime := time.Now().Add(-24 * time.Hour)

	// Gunakan fungsi time_bucket dari TimescaleDB untuk agregasi
	// Tanda kutip ganda (" ") penting untuk nama kolom yang case-sensitive
	query := fmt.Sprintf(`
        SELECT 
            time_bucket(INTERVAL '1 %s', "Timestamp") AS bucket,
            AVG("%s") as avg_value
        FROM readings
        WHERE "DeviceId" = $1 AND "Timestamp" > $2
        GROUP BY bucket
        ORDER BY bucket;
    `, interval, parameter) // Hati-hati dengan string formatting, tapi di sini aman karena input sudah divalidasi

	rows, err := pool.Query(context.Background(), query, deviceID, startTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []HistoricalPoint
	for rows.Next() {
		var point HistoricalPoint
		if err := rows.Scan(&point.TimeBucket, &point.AvgValue); err != nil {
			return nil, err
		}
		results = append(results, point)
	}

	return results, nil
}
