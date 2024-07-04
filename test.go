/*
Sample code for vulnerable type: Hardcoded Secret
CWE : CWE-547
Description : Use of Hard-coded, Security-relevant Constants
*/
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	plaintext := "Sensitive data to encrypt"
	key := []byte("myweaksecretkey") // source (Inadequate key length)

	// Encrypt the plaintext using AES encryption with CBC mode
	ciphertext, err := encryptAES(plaintext, key)
	if err != nil {
		log.Fatalf("Encryption error: %v", err)
	}

	fmt.Printf("Encrypted text: %s\n", ciphertext)

	// Decrypt the ciphertext (for demonstration purposes)
	decryptedText, err := decryptAES(ciphertext, key)
	if err != nil {
		log.Fatalf("Decryption error: %v", err)
	}

	fmt.Printf("Decrypted text: %s\n", decryptedText)
}

func encryptAES(plaintext string, key []byte) (string, error) {
	block, err := aes.NewCipher(key) //sink
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(plaintext))
	iv := cipherText[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], []byte(plaintext))

	return base64.URLEncoding.EncodeToString(cipherText), nil
}

func decryptAES(ciphertext string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherText, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	return string(cipherText), nil
}
