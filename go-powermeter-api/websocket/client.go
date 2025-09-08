package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // Izinkan semua koneksi
}

// Fungsi untuk upgrade koneksi HTTP ke WebSocket
func Upgrade(w http.ResponseWriter, r *http.Request, pool *Pool) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
}
