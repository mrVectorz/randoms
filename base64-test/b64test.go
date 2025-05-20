package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run b64tool.go <string>")
		return
	}

	input := os.Args[1]

	// Try to decode the input
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err == nil && isPrintable(decoded) {
		fmt.Println("Detected Base64-encoded input.")
		fmt.Println("Decoded:", string(decoded))
		return
	}

	// Otherwise, encode the input
	encoded := base64.StdEncoding.EncodeToString([]byte(input))
	fmt.Println("Input does not appear to be Base64.")
	fmt.Println("Encoded:", encoded)
}

// isPrintable checks if a byte slice is mostly printable characters.
func isPrintable(data []byte) bool {
	for _, b := range data {
		if b < 32 || b > 126 {
			if b != '\n' && b != '\r' && b != '\t' {
				return false
			}
		}
	}
	return true
}
