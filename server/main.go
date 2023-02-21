package main

import (
	"github.com/fatih/color"
	"net"
)

func main() {
	ip := "0.0.0.0"
	port := 8858
	addr := &net.UDPAddr{
		IP:   net.ParseIP(ip), //\
		Port: port,            // | These are the default values
	}
	conn, err := net.ListenUDP("udp", addr)
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 1024)
	color.Green("Listening on %s:%d", ip, port)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}
		color.Green("Received %s from %s", string(buffer[0:n]), addr.String())
		_, err = conn.WriteToUDP([]byte("Hello "+string(buffer[0:n])), addr)
		if err != nil {
			panic(err)
		}
	}
}
