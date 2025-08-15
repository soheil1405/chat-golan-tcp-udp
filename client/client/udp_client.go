package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func StartUDPClient(address string) error {
	conn, err := net.Dial("udp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	serverReader := bufio.NewReader(conn)
	consoleReader := bufio.NewReader(os.Stdin)

	// Ask username
	fmt.Print("Enter your name: ")
	username, _ := consoleReader.ReadString('\n')
	username = strings.TrimSpace(username)
	conn.Write([]byte(username)) // first message is username

	// Read welcome message from server (optional)
	reply, _ := serverReader.ReadString('\n')
	fmt.Print(reply)

	// Chat loop
	for {
		input, _ := consoleReader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input != "" {
			conn.Write([]byte(input))
		}
	}

	return nil
}

func main() {
	StartUDPClient("127.0.0.1:9001")
}
