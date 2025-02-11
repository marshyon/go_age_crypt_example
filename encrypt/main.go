package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"filippo.io/age"
	"github.com/joho/godotenv"
)

func main() {

	plaintextFile := "input.txt"
	encryptedFile := "encrypted_file.age"

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	publicKey := os.Getenv("PUBLIC_KEY")
	if publicKey == "" {
		log.Fatalf("PUBLIC_KEY not set in .env file")
	}

	recipient, err := age.ParseX25519Recipient(publicKey)
	if err != nil {
		log.Fatalf("Failed to parse public key %q: %v", publicKey, err)
	}

	inFile, err := os.Open(plaintextFile)
	if err != nil {
		log.Fatalf("Failed to open input file: %v", err)
	}
	defer inFile.Close()

	out, err := os.Create(encryptedFile)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer out.Close()

	w, err := age.Encrypt(out, recipient)
	if err != nil {
		log.Fatalf("Failed to create encrypted file: %v", err)
	}

	if _, err := io.Copy(w, inFile); err != nil {
		log.Fatalf("Failed to write to encrypted file: %v", err)
	}

	if err := w.Close(); err != nil {
		log.Fatalf("Failed to close encrypted file: %v", err)
	}

	fileInfo, err := out.Stat()
	if err != nil {
		log.Fatalf("Failed to get file info: %v", err)
	}

	fmt.Printf("Encrypted file size: %d\n", fileInfo.Size())

}
