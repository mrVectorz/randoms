package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {
	// Create a raw socket using AF_PACKET
	sock, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, syscall.ETH_P_ALL)
	if err != nil {
		log.Fatalf("Failed to create socket: %v", err)
	}
	sockb, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, syscall.ETH_P_ALL)

	fmt.Println("AF_PACKET socket created successfully!")

	defer syscall.Close(sock)
	defer syscall.Close(sockb)
}
