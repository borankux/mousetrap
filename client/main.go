package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	ip := "0.0.0.0"
	port := 8858
	addr := &net.UDPAddr{
		IP:   net.ParseIP(ip),
		Port: port,
	}
	conn, err := net.DialUDP("udp", nil, addr)
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, 1024)
	for {
		count, err := conn.Write([]byte("Mirzat"))
		if err != nil {
			panic(err)
		}
		fmt.Println("Sent ", count, " bytes")
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Println("Received ", string(buffer[0:n]), " from ", addr)
		time.Sleep(time.Second * 1)
	}
}
