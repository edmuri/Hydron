package main

import (
	"fmt"
	"net"
)

type PacketJob struct {
	src net.IP
	dst net.IP
	msg []byte
	amt int
}

// TODO: Adding more argument handling
func verifyArgs(args []string) bool {
	fmt.Printf("args")
	return true
}
