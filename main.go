package main

import (
	"fmt"
	"net"
	"os"
)

func setup_packet() {
	ip := "127.0.0.1"
	conn, err := net.Dial("udp", ip)

	if err != nil {
		fmt.Printf("[!] Failed to connected to address: %v\n", err)
	}
	defer conn.Close()

	payload := []byte("Hello")

	_, err = conn.Write(payload)

	if err != nil {
		fmt.Println("[!] Error sending packet:", err)
		return
	}
	fmt.Println("[-] Packet sent successfully!")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("[!] Not enough arguments")
		return
	}

	args := os.Args[1:]

	status := verifyArgs(args)

	fmt.Println(status)
	fmt.Println(args)
	setup_packet()
}
