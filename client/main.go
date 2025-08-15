// client/main.go
package main

import (
	"chat-app/client/client"
	"fmt"
)

func main() {
	var (
		choice int
	)

	fmt.Println("Choose connection type:")
	fmt.Println("1 - TCP")
	fmt.Println("2 - UDP")
	fmt.Scanln(&choice)

	if choice == 1 {
		client.StartTCPClient("localhost:9000")
	} else {
		client.StartUDPClient("localhost:9001")
	}
}
