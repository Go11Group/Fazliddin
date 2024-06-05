package main

import (
	"bufio"
	"fmt"
	"net"
)

var mp = make(map[net.Conn]bool)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr().String())

	mp[conn] = true
	reader := bufio.NewReader(conn)
	for {

		fmt.Println(mp)
		// Read data from the connection.
		message, err := reader.ReadString('\n')
		if err != nil {
			delete(mp, conn)
			fmt.Println("Error reading message:", err)
			return
		}
		fmt.Print("Received message: ", string(message))
		message = message[:len(message)-1]
		for i := range mp {
			fmt.Println("sende_________")
			_, err := i.Write([]byte(message + "\n"))
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}