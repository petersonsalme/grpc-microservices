package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/petersonsalme/go-grpc-project/go-grpc-auth-svc/internal/models"
)

type Wrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type jwtClaims struct {
	jwt.RegisteredClaims
	ID    int64
	Email string
}

func (w *Wrapper) GenerateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtClaims{
		ID:    user.ID,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(w.ExpirationHours))),
			Issuer:    w.Issuer,
		},
	})

	signedToken, err := token.SignedString([]byte(w.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (w *Wrapper) ValidateToken(signedToken string) (*jwtClaims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(w.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now().Local()) {
		return nil, errors.New("JWT is expired")
	}

	return claims, nil
}
