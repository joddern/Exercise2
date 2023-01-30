package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen to incoming connections on port 30000
	addr, err := net.ResolveUDPAddr("udp", ":30000")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Receive data from the server
	buf := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Received", n, "bytes:", string(buf[:n])) //change from buf[:n] to just buf
	}
}
