package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Wrong Input! Usage:", os.Args[0], "127.0.0.1:1234", "ABC")
		return
	}

	dest := os.Args[1]
	msg := os.Args[2]

	conn, err := net.Dial("udp", dest)
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	fmt.Fprintf(conn, msg)
	//---- if you want a response
	//p := make([]byte, 2048)
	//_, err = bufio.NewReader(conn).Read(p)
	//if err != nil {
	//	fmt.Printf("Some error %v\n", err)
	//}
	//fmt.Println("before CC")
	conn.Close()
}
