package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: decrypt <input_file> <output_file>")
		return
	}

	if len(os.Args) == 3 {
		decrypt(os.Args[1], os.Args[2])
	} else {
		decrypt(os.Args[1], "")
	}
}

func decrypt(inputFilePath, outputFilePath string) {
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

	ciphertext, err := os.ReadFile(inputFilePath)
	if err != nil {
		fmt.Println("Error reading encrypted file:", err)
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

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println("Invalid ciphertext (too short)")
		return
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}

	if outputFilePath == "" {
		fmt.Println(string(plaintext))
		return
	}

	err = os.WriteFile(outputFilePath, plaintext, 0644)
	if err != nil {
		fmt.Println("Error writing decrypted file:", err)
		return
	}

	fmt.Println("Decryption successful.")
}
