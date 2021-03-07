package token

import (
	"github.com/dgrijalva/jwt-go"
)

// CreateToken lib uses jwt-go to sign an integer with ES256 encryption
func CreateToken(userID uint64, secretKey string) (string, error) {
	jtClaims := jwt.MapClaims{}

	jtClaims["userId"] = userID
	jtClaims["authorized"] = true
	signer := jwt.NewWithClaims(jwt.SigningMethodHS256, jtClaims)
	token, err := signer.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}
