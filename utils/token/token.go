package token

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type JwtCustomClaims struct {
	Username string `json:"user_id"`
	jwt.StandardClaims
}

// generate JWT.
func GenerateJWT(ctx context.Context, tokenLifeTime time.Duration, secretKey string, username string) (string, error) {
	if secretKey == "" {
		return "", fmt.Errorf("secret key not found")
	}

	claims := &JwtCustomClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * tokenLifeTime).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ToHashPassword hashes the password using bcrypt
func ToHashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}

// VerifyToken is check the validity of the token and return contents.
func VerifyToken(token string, secretKey string) (*JwtCustomClaims, error) {
	if token == "" {
		return nil, fmt.Errorf("authorize token required")
	}
	if secretKey == "" {
		return nil, fmt.Errorf("secret key not found")
	}

	claims := JwtCustomClaims{}
	_, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		if e, ok := err.(*jwt.ValidationError); ok {
			if e.Errors == jwt.ValidationErrorExpired {
				return nil, err
			}
			return nil, fmt.Errorf("token invalid")
		}
		return nil, err
	}

	return &claims, nil
}
