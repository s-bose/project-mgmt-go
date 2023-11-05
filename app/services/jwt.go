package services

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type ClaimsPayload struct {
	sub string
	jwt.StandardClaims
}

func CreateAccessToken(payload string) (map[string]string, error) {
	claims := &ClaimsPayload{
		sub: payload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &ClaimsPayload{
		sub: payload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24*30)).Unix(),
		},
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return nil, err
	}
	refreshTokenString, err := jwt.NewWithClaims(jwt.SigningMethodES256, refreshClaims).SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  tokenString,
		"refresh_token": refreshTokenString,
	}, nil
}

func ValidateJwtToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

	return true, nil
}
