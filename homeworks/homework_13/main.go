package main

import (
	"fmt"
	"net"
	"sync"
)

var c = make(map[string]net.Addr)
var m = &sync.Mutex{}

func main() {
	conn, err := net.ListenPacket("udp", ":8080")
	if err != nil {
		fmt.Println("Tinglashni sozlashda xato:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Server 8080-portda tinglamoqa...")

	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			fmt.Println("Xabarni o'qishda xato:", err)
			continue
		}

		message := string(buffer[:n])
		fmt.Printf("Qabul qilingan xabar: %s\n", message)
		addClient(addr)
		broadcastMessage(conn, addr, message)
	}
}

func addClient(addr net.Addr) {
	m.Lock()
	defer m.Unlock()
	c[addr.String()] = addr
}

func broadcastMessage(conn net.PacketConn, senderAddr net.Addr, message string) {
	m.Lock()
	defer m.Unlock()
	for _, clientAddr := range c {
		//if clientAddr.String() != senderAddr.String() {
		response := fmt.Sprintf("%s => %s", senderAddr.String(), message)
		_, err := conn.WriteTo([]byte(response), clientAddr)
		if err != nil {
			fmt.Println("Xabarni yozishda xato:", err)
		}
		//}
	}
}
