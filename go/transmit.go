package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("testing")

	// # Sending UDP packets

	// Make a connection, a net.Conn object
	connection, err := net.Dial("udp", "10.30.42.227:20000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer connection.Close()

	// Send messages on this connection to the server (as goroutine so that we can )
	go func() {
		for i := 0; i < 10; i++ {
			message := fmt.Sprintf("Hello, server! Msg#: %d", i) // either just make it a slice of bytes immediately, or like below:

			_, err = connection.Write([]byte(message))
			if err != nil {
				fmt.Println(err)
				return // strict no?
			}
			fmt.Println("Message sent:", message)
			time.Sleep(time.Second)
		}
	}()

	// Receive the responses
	connection_2, err := net.ListenPacket("udp", ":20001")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer connection_2.Close()
	for {
		buf := make([]byte, 1024)
		n, addr, err := connection_2.ReadFrom(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Output received response
		fmt.Println("Received response from address:", addr, ". Number of bytes:", n, ". Message:", string(buf[:n]))
	}

}
