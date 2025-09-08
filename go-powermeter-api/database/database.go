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
		"Timestamp", "DeviceId", "EnergyKWh", "CurrentL1", "CurrentL2", "CurrentL3", "CurrentAverage",
		"VoltageL1ToL2", "VoltageL1ToL3", "VoltageL2ToL3", "Voltage3PhaseAverage", "VoltageL1ToN",
		"VoltageL2ToN", "VoltageL3ToN", "Voltage1PhaseAverage", "ActivePowerL1", "ActivePowerL2",
		"ActivePowerL3", "ActivePowerTotal", "ReactivePowerL1", "ReactivePowerL2", "ReactivePowerL3",
		"ReactivePowerTotal", "PowerFactorL1", "PowerFactorL2", "PowerFactorL3", "PowerFactorTotal",
		"HarmonicDistortionCurrent", "HarmonicDistortionVoltage3Ph", "HarmonicDistortionVoltage1Ph",
		"Frequency",
	}

	rows := make([][]interface{}, len(readings))
	for i, r := range readings {
		// Urutan data HARUS sama persis dengan daftar kolom di atas
		rows[i] = []interface{}{
			r.Timestamp, r.DeviceID, r.EnergyKWh, r.CurrentL1, r.CurrentL2, r.CurrentL3, r.CurrentAverage,
			r.VoltageL1ToL2, r.VoltageL1ToL3, r.VoltageL2ToL3, r.Voltage3PhaseAverage, r.VoltageL1ToN,
			r.VoltageL2ToN, r.VoltageL3ToN, r.Voltage1PhaseAverage, r.ActivePowerL1, r.ActivePowerL2,
			r.ActivePowerL3, r.ActivePowerTotal, r.ReactivePowerL1, r.ReactivePowerL2, r.ReactivePowerL3,
			r.ReactivePowerTotal, r.PowerFactorL1, r.PowerFactorL2, r.PowerFactorL3, r.PowerFactorTotal,
			r.HarmonicDistortionCurrent, r.HarmonicDistortionVoltage3Ph, r.HarmonicDistortionVoltage1Ph,
			r.Frequency,
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
