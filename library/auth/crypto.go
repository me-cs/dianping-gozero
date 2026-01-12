package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
)

// encryptionKey is the AES-256 key (32 bytes)
// IMPORTANT: In production, this should be loaded from environment variables or secure configuration
// Generated key: ba4c52902510af0df755679bca7a3cfaa75c47b5f283a59fc3eb3cd10a464a43
var encryptionKey = []byte("dianping-gozero-aes256-key-2026!")

// SetEncryptionKey allows setting the encryption key from configuration
func SetEncryptionKey(key []byte) error {
	if len(key) != 32 {
		return errors.New("encryption key must be exactly 32 bytes for AES-256")
	}
	encryptionKey = key
	return nil
}

// EncodeUser encrypts user data using AES-GCM and returns hex-encoded string
// Format: hex(nonce + ciphertext)
func EncodeUser(u User) (string, error) {
	// Marshal user to JSON
	data, err := json.Marshal(u)
	if err != nil {
		return "", err
	}

	// Create AES cipher
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}

	// Create GCM mode
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Generate random nonce (12 bytes for GCM)
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Encrypt data
	ciphertext := aesGCM.Seal(nil, nonce, data, nil)

	// Prepend nonce to ciphertext and encode as hex
	result := append(nonce, ciphertext...)
	return hex.EncodeToString(result), nil
}

// DecodeUser decrypts hex-encoded encrypted user data using AES-GCM
func DecodeUser(s string) (User, error) {
	// Decode hex string
	data, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}

	// Create AES cipher
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, err
	}

	// Create GCM mode
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Check minimum length
	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	// Extract nonce and ciphertext
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	// Decrypt data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON
	u := &user{}
	err = json.Unmarshal(plaintext, u)
	return u, err
}
