package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type Payload map[string]interface{}

func (p Payload) SignWithHS256(secret string) (string, error) {
	claims := jwt.MapClaims(p)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func VerifyWithHS256(secret string, tokenStr string) (Payload, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signature method: %v", t.Method.Alg())
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("token is invalid: %v", err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		p := Payload(claims)
		return p, nil
	} else {
		return nil, fmt.Errorf("token is invalid")
	}
}
