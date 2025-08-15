package main

import (
	"chat-app/server/server"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		server.StartTCPServer("9000")
		wg.Done()
	}()
	go func() {
		server.StartUDPServer("9001")
		wg.Done()
	}()

	wg.Wait()
}
