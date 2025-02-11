package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"filippo.io/age"
	"github.com/joho/godotenv"
)

// DO NOT hardcode the private key. Store it in a secret storage solution,
// on disk if the local machine is trusted, or have the user provide it.

func main() {

	encryptedFile := "encrypted_file.age"
	decryptedFile := "decrypted_file.txt"

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	privateKey := os.Getenv("PRIVATE_KEY")
	if privateKey == "" {
		log.Fatalf("PUBLIC_KEY not set in .env file")
	}

	identity, err := age.ParseX25519Identity(privateKey)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	f, err := os.Open(encryptedFile)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	r, err := age.Decrypt(f, identity)
	if err != nil {
		log.Fatalf("Failed to open encrypted file: %v", err)
	}

	outFile, err := os.Create(decryptedFile)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer outFile.Close()

	if _, err := io.Copy(outFile, r); err != nil {
		log.Fatalf("Failed to write decrypted data to file: %v", err)
	}

	fmt.Println("Decryption successful, data written to", outFile.Name())

}
