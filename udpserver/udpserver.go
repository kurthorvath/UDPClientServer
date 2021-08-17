package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"gopkg.in/natefinch/lumberjack.v2"
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("From server: Hello I got your message "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func httpServ() {
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.ListenAndServe(":8888", nil)
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Wrong Input! Usage:", os.Args[0], "1234")
		return
	}

	dest := os.Args[1]
	port, _ := strconv.Atoi(dest)

	log.SetOutput(&lumberjack.Logger{
		Filename:   "server.html",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})

	go httpServ()

	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("127.0.0.1"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	for {
		p := make([]byte, 2048)
		_, remoteaddr, err := ser.ReadFromUDP(p)

		p = bytes.Trim(p, "\x00")
		plen := len(p)
		fmt.Println("Read: ", remoteaddr, " size: [", plen, "] ", p, string(p))
		log.Println("Read: ", remoteaddr, " size: [", plen, "] ", p, string(p), "<br>")
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		//in case we want to send an answer
		//go sendResponse(ser, remoteaddr)
	}

}
