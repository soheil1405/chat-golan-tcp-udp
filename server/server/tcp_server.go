package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

var (
	clientNames = make(map[net.Conn]string)
	tcpmutex    sync.Mutex
)

// Broadcast to all clients except sender (or all if sender is nil)
func broadcast(sender net.Conn, message string) {
	tcpmutex.Lock()
	defer tcpmutex.Unlock()
	for conn := range clientNames {
		if conn != sender {
			conn.Write([]byte(message))
		}
	}
}

func StartTCPServer(port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	fmt.Println("TCP server running on port", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err)
			continue
		}
		go handleTCPConnection(conn)
	}
}

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read name:", err)
		return
	}
	name = strings.TrimSpace(name)

	// Save the name
	tcpmutex.Lock()
	clientNames[conn] = name
	tcpmutex.Unlock()

	// Notify everyone that a new client joined
	broadcast(nil, fmt.Sprintf("[SERVER] %s joined the chat\n", name))
	// Send welcome message to the new client
	conn.Write([]byte(fmt.Sprintf("Welcome, %s!\n", name)))

	// Read messages from client
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("[SERVER] %s disconnected\n", name)
			tcpmutex.Lock()
			delete(clientNames, conn)
			tcpmutex.Unlock()
			broadcast(nil, fmt.Sprintf("[SERVER] %s left the chat\n", name))
			return
		}
		message = strings.TrimSpace(message)
		fmt.Printf("[%s] %s\n", name, message)
		broadcast(conn, fmt.Sprintf("[%s] %s\n", name, message))
	}
}
