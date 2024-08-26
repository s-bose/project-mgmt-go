package auth

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type ClaimsPayload struct {
	UserID string
	jwt.StandardClaims
}

func CreateJWTToken(UserID string) (map[string]string, error) {
	claims := &ClaimsPayload{
		UserID: UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(time.Hour * time.Duration(24)).Unix(),
			Subject:   "access_token",
		},
	}

	refreshClaims := &ClaimsPayload{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(time.Hour * time.Duration(24*30)).Unix(),
			Subject:   "refresh_token",
		},
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return nil, err
	}
	refreshTokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  tokenString,
		"refresh_token": refreshTokenString,
	}, nil
}

func ValidateJWTToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if !token.Valid {
		log.Println("invalid token")
		return nil, errors.New("invalid token")
	}

	claims := token.Claims.(*ClaimsPayload)
	if claims.ExpiresAt < time.Now().UTC().Unix() {
		log.Println("token expired")
		return nil, errors.New("token expired")
	}

	return token, err
}
