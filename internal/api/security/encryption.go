package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/blowfish"
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

func BlowfishEncrypt(ppt []byte) ([]byte, error) {
	key := []byte(os.Getenv("BLOWFISH_SECRET_KEY"))
	ecipher, err := blowfish.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, blowfish.BlockSize+len(ppt))
	eiv := ciphertext[:blowfish.BlockSize]
	ecbc := cipher.NewCBCEncrypter(ecipher, eiv)
	ecbc.CryptBlocks(ciphertext[blowfish.BlockSize:], ppt)
	return ciphertext, nil
}

func BlowfishChecksizeAndPad(pt []byte) []byte {
	modulus := len(pt) % blowfish.BlockSize
	if modulus != 0 {
		padlen := blowfish.BlockSize - modulus
		for i := 0; i < padlen; i++ {
			pt = append(pt, 0)
		}
	}
	return pt
}
