package main

import(
	"fmt"
	"net"
	"bufio"
	_"strings"
)

func main(){
	listner, err := net.Listen("tcp", ":8080")
	if err != nil{
		fmt.Println("ERROR: ", err)
		return
	}

	defer listner.Close()
	fmt.Println("Server is listenin on port: 8080... ;")

	for{
		conn, err := listner.Accept()
		if err != nil{
			fmt.Println("ERROR: ", err)
			return
		}
		go hendleConn(conn)
	}
}

var ls = [][]byte{}

func hendleConn(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for{
		messege, err := reader.ReadString('\n')
		if err != nil{
			fmt.Println("ERROR: ", err)
			return
		}
		ls = append(ls, messege)
		fmt.Println(messege)
		_, err = conn.Write([]byte(messege))
		if err != nil{
			fmt.Println("ERROR: ", err)
		}
	}
	fmt.Println("hello")
}
