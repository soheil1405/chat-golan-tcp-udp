package server

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

var (
	udpClients = make(map[string]string) // addr.String() -> username
	udpMutex   sync.Mutex
)

func StartUDPServer(port string) {
	addr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("UDP server running on port", port)

	buffer := make([]byte, 1024)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading UDP:", err)
			continue
		}

		message := strings.TrimSpace(string(buffer[:n]))

		udpMutex.Lock() // use udpMutex here
		username, ok := udpClients[remoteAddr.String()]
		if !ok {
			// First message is username
			udpClients[remoteAddr.String()] = message
			username = message
			fmt.Printf("[UDP] %s joined the chat\n", username)
			conn.WriteToUDP([]byte(fmt.Sprintf("Welcome, %s!\n", username)), remoteAddr)
			udpMutex.Unlock()
			continue
		}
		udpMutex.Unlock()

		fmt.Printf("[UDP][%s] %s\n", username, message)
	}
}
