package websocket

import (
	"log"
	"github.com/gorilla/websocket"
)

// Client merepresentasikan satu koneksi WebSocket.
type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

// Read akan terus mendengarkan pesan dari koneksi klien.
// Ini penting untuk menjaga koneksi dan mendeteksi saat klien terputus.
func (c *Client) Read() {
	defer func() {
		// Saat loop ini berhenti (karena klien disconnect),
		// kirim sinyal unregister ke pool.
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		// Baca pesan dari koneksi. Jika ada error (misal, klien disconnect),
		// loop akan berhenti secara otomatis.
		if _, _, err := c.Conn.ReadMessage(); err != nil {
			log.Println("Klien terputus:", err)
			break
		}
	}
}