package secret

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Encrypt encrypts the plain text with bcrypt.
func Encrypt(source string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashed), err
}

// Compare compares the encrypted text with the plain text if it's the same.
func Compare(hashed, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}

// Sign issue a jwt token based on secret ID, secret Key, iss and aud.
func Sign(id string, key string, iss, aud string) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Minute).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Add(0).Unix(),
		"aud": aud,
		"iss": iss,
	}

	// create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["kid"] = id

	// Sign the token with the specified secret.
	return token.SignedString([]byte(key))
}
