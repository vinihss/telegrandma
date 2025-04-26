package deamon

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Allow all origins (for development only - restrict in production)
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketStart() {
	// Serve static files (optional)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// WebSocket endpoint
	http.HandleFunc("/ws", handleWebSocket)

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade to WebSocket: %v", err)
		return
	}
	defer conn.Close()

	log.Printf("New WebSocket connection from %s", r.RemoteAddr)

	// Communication loop
	for {
		// Read message from client
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			break
		}

		log.Printf("Received: %s", message)

		// Echo the message back to the client
		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Printf("Write error: %v", err)
			break
		}
	}
}
