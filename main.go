package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	// 1. master.key
	masterKey := generateRandomHex(32)
	masterKeyBase64 := base64.StdEncoding.EncodeToString([]byte(masterKey))
	fmt.Printf("master.key: %s\n", masterKeyBase64)

	// 2. Генерируем secret.key
	secretKey := generateRandomHex(32)
	secretKeyBase64 := base64.StdEncoding.EncodeToString([]byte(secretKey))
	fmt.Printf("secret.key: %s\n", secretKeyBase64)

	// 3. Генерируем hudson.util.Secret
	hudsonSecret := generateHudsonSecret([]byte(masterKey))
	hudsonSecretBase64 := base64.StdEncoding.EncodeToString([]byte(hudsonSecret))
	fmt.Printf("hudson.util.Secret: %s\n", hudsonSecretBase64)

	// 4. Генерируем com.cloudbees.plugins.credentials.SecretBytes.KEY
	secretBytesKey := generateAESKey()
	secretBytesKeyBase64 := base64.StdEncoding.EncodeToString([]byte(secretBytesKey))
	fmt.Printf("com.cloudbees.plugins.credentials.SecretBytes.KEY: %s\n", secretBytesKeyBase64)
}

func generateRandomHex(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

func generateHudsonSecret(masterKey []byte) []byte {
	hasher := sha256.New()
	hasher.Write(masterKey)
	return hasher.Sum(nil)
}

func generateAESKey() []byte {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	return key
}
