package main

import (
	"flag"
	"fmt"
	"net"
)

func read(conn net.Conn) {
	//TODO In a continuous loop, read a message from the server and display it.
	defer conn.Close()
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error", err)
			return
		}
		fmt.Println("server:", string(buffer[:n]))

	}
}

func write(conn net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
	defer conn.Close()
	for {
		var message string
		fmt.Println("You: ")
		_, err := fmt.Scanln(&message)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		_, err = conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Error sending input:", err)
			return
		}

	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	//TODO Try to connect to the server
	conn, err := net.Dial("tcp", *addrPtr)
	if err != nil {
		fmt.Println("Error connecting o server:", err)
		return
	}
	fmt.Println("Connected to server at", *addrPtr)

	//TODO Start asynchronously reading and displaying messages
	go read(conn)
	//TODO Start getting and sending user messages.
	write(conn)
}
