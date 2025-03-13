package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	const maxSockets = 500
	const batchSize = 50

	var sockets []int

	for len(sockets) < maxSockets {
		for i := 0; i < batchSize && len(sockets) < maxSockets; i++ {
			sock, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, syscall.ETH_P_ALL)
			if err != nil {
				log.Fatalf("Failed to create socket: %v", err)
			}
			sockets = append(sockets, sock)
		}
		fmt.Printf("Opened %d sockets so far...\n", len(sockets))
		time.Sleep(5 * time.Second)
	}

	fmt.Println("All sockets created successfully!")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.Signal(7))
	<-sigChan

	// Close all sockets before exiting
	for _, sock := range sockets {
		syscall.Close(sock)
	}

	fmt.Println("All sockets closed successfully!")
}
