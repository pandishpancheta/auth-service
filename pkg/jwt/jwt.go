package jwt

import (
	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtWrapper struct {
	SecretKey       string
	ExpirationHours int64
}

func (j *JwtWrapper) GenerateToken(id, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JwtWrapper) ValidateToken(tokenString string) (*uuid.UUID, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, jwt.ErrSignatureInvalid
	}

	id := claims["id"].(string)
	uuid, err := uuid.FromString(id)
	if err != nil {
		return nil, err
	}

	return &uuid, err
}
