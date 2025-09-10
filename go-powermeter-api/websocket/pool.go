package websocket

import (
	"encoding/json"
	"fmt"

	"github.com/amaulanah/powermeterapi/models" // Ganti dengan path modul Anda
)

// Pool untuk mengelola semua koneksi client
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

// Method untuk menjalankan pool di background
func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Koneksi baru. Ukuran pool:", len(pool.Clients))
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Koneksi terputus. Ukuran pool:", len(pool.Clients))
		case message := <-pool.Broadcast:
			// 2. TAMBAHKAN PRINT DEBUG DI SINI
			// =======================================================
			jsonData, _ := json.Marshal(message)
			fmt.Println("Mengirim JSON:", string(jsonData))
			// Kirim pesan ke semua client yang terhubung
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					continue
				}
			}
		}
	}
}
