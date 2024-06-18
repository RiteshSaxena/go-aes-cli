package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: encrypt <input_file> <output_file>")
		return
	}

	encrypt(os.Args[1], os.Args[2])
}

func encrypt(inputFilePath, outputFilePath string) {
	keyHex, err := os.ReadFile("keyFile.txt")
	if err != nil {
		fmt.Println("Error reading key file:", err)
		return
	}

	key, err := hex.DecodeString(string(keyHex))
	if err != nil {
		fmt.Println("Error decoding key:", err)
		return
	}

	plaintext, err := os.ReadFile(inputFilePath)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error creating cipher block:", err)
		return
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("Error creating GCM:", err)
		return
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println("Error generating nonce:", err)
		return
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)

	err = os.WriteFile(outputFilePath, ciphertext, 0644)
	if err != nil {
		fmt.Println("Error writing encrypted file:", err)
		return
	}

	fmt.Printf("Encryption successful.")
}
