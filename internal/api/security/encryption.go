package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"log"
	"os"
)

var (
	errDataEncryption = errors.New("Error on data encryption")
)

func AesEncrypt(text []byte) ([]byte, error) {
	aesSecretKey := []byte(os.Getenv("AES_SECRET_KEY"))
	c, err := aes.NewCipher([]byte(aesSecretKey))
	if err != nil {
		log.Printf("Cannot create the cipher: %v\n", err)
		return nil, errDataEncryption
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Printf("Error generating GCM: %v\n", err)
		return nil, errDataEncryption
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Printf("Error on populate nonce: %v\n", err)
		return nil, errDataEncryption
	}
	return gcm.Seal(nonce, nonce, text, nil), nil
}
