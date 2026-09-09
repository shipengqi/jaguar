package secret

import (
	"golang.org/x/crypto/bcrypt"
)

// Encrypt encrypts plain text with bcrypt.
func Encrypt(plain string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(hashed), err
}

// Compare compares bcrypt-encrypted hash with plain text.
func Compare(hashed, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
}
