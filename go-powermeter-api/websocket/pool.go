package websocket

import (
	"fmt"

	"github.com/amaulanah/powermeterapi/models" // Sesuaikan path modul Anda
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan []models.PowerMeterReading
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []models.PowerMeterReading),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Koneksi baru. Ukuran pool:", len(pool.Clients))

		case client := <-pool.Unregister:
			if _, ok := pool.Clients[client]; ok {
				delete(pool.Clients, client)
				fmt.Println("Koneksi terputus. Ukuran pool:", len(pool.Clients))
			}

		case message := <-pool.Broadcast:
			fmt.Println("Menyiarkan pesan ke semua klien...")
			// Kirim pesan ke semua klien yang terhubung
			for client := range pool.Clients {
				// Perbaikan Kunci: Gunakan goroutine untuk mengirim pesan
				// agar tidak memblokir loop utama jika satu klien lambat/error.
				go func(c *Client) {
					if err := c.Conn.WriteJSON(message); err != nil {
						fmt.Println("Gagal mengirim ke klien, akan dihapus:", err)
						// Jika gagal mengirim, kirim klien ini ke channel Unregister
						c.Pool.Unregister <- c
						c.Conn.Close()
					}
				}(client) // <-- Jalankan fungsi anonim ini sebagai goroutine
			}
		}
	}
}