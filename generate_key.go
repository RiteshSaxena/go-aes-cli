package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	key := make([]byte, 32) // AES-256 requires a 32-byte key
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		fmt.Println("Error generating key:", err)
		return
	}

	fmt.Println("Key:", hex.EncodeToString(key))

	keyFileName := "keyFile.txt"
	keyFile, err := os.Create(keyFileName)
	if err != nil {
		fmt.Println("Error creating key file:", err)
		return
	}
	defer keyFile.Close()
	keyFile.WriteString(hex.EncodeToString(key))
}
