package secret

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

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
