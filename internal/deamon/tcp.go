package deamon

import (
	"bufio"
	"log"
	"net"
	"strings"
)

type TCPServer struct {
	Port int
}

func (t TCPServer) New() TCPServer {
	t.Port = 8080
	return t
}

func (t TCPServer) Start() {
	// Start listening on port 8080
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()

	log.Println("TCP server started on :8080")

	for {
		// Wait for and accept new connections
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		// Handle each connection in a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	clientAddr := conn.RemoteAddr().String()
	log.Printf("Client connected: %s", clientAddr)

	reader := bufio.NewReader(conn)

	for {
		// Read client message
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Client %s disconnected: %v", clientAddr, err)
			return
		}

		// Process message (trim newline and convert to uppercase for example)
		message = strings.TrimSpace(message)
		response := strings.ToUpper(message)

		// Send response back to client
		_, err = conn.Write([]byte(response + "\n"))
		if err != nil {
			log.Printf("Error writing to client %s: %v", clientAddr, err)
			return
		}

		log.Printf("Client %s: %s -> %s", clientAddr, message, response)
	}
}
