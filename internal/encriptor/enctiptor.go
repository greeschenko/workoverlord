package encriptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	//"encoding/hex"
	"io"
	"log"
)

type Encriptor struct {}

func NewEncriptor() *Encriptor {
	return &Encriptor{}
}

func (e Encriptor) Encrypt(data []byte, SECRETKEY [32]byte) []byte {

	// generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher(SECRETKEY[:])
	// if there are any errors, handle them
	if err != nil {
		log.Fatalln(err)
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	// if any error generating new GCM
	// handle them
	if err != nil {
		log.Fatalln(err)
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalln(err)
	}

	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.

	res := gcm.Seal(nonce, nonce, data, nil)

	return res
}

func (e Encriptor) Descrypt(data []byte, SECRETKEY [32]byte) ([]byte, error) {

	ciphertext := data

	c, err := aes.NewCipher(SECRETKEY[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, err
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	res := plaintext
	return res, nil
}
