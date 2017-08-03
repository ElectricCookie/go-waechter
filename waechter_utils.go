package waechter

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/scrypt"
)

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateRandomString(length int) string {
	b, err := generateRandomBytes(length)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}

func hash(input string, salt string) string {
	b, err := scrypt.Key([]byte(input), []byte(salt), 16384, 8, 1, 32)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}
