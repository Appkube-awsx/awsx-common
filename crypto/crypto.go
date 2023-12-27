package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/Appkube-awsx/awsx-common/config"
	"io"
)

func Encrypt(plainText string) (string, error) {
	data := []byte(plainText)
	block, err := aes.NewCipher(config.Key)
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not necessarily secure.
	// In this example, we use a random 16-byte IV for simplicity.
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(data))
	stream.XORKeyStream(ciphertext, data)

	// Prepend the IV to the ciphertext.
	ciphertext = append(iv, ciphertext...)
	encoded := base64.StdEncoding.EncodeToString(ciphertext)

	return encoded, nil
}

func Decrypt(encodedText string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encodedText)
	if err != nil {
		fmt.Println("Base64 decoding error:", err)
		return "", err
	}
	block, err := aes.NewCipher(config.Key)
	if err != nil {
		return "", err
	}

	// Extract the IV from the beginning of the ciphertext.
	iv := data[:aes.BlockSize]
	ciphertext := data[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}
