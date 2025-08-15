package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func StartTCPClient(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	go func() {
		serverReader := bufio.NewReader(conn)
		for {
			reply, err := serverReader.ReadString('\n')
			if err != nil {
				fmt.Println("Disconnected from server")
				os.Exit(0)
			}
			fmt.Print(reply)
		}
	}()

	consoleReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := consoleReader.ReadString('\n')
		conn.Write([]byte(input))
	}
}
