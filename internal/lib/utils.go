package lib

import (
	"errors"
	"fmt"
	"time"

	"github.com/RE-PIT-BEM/re-pit-backend/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func CreateJWTToken(data jwt.Claims) (string, error) {
	viper := config.Load()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	tokenString, err := token.SignedString([]byte(viper.GetString("secret.jwt-secret")))

	return tokenString, err
}

func ParseJWTToken(tokenString string) (jwt.MapClaims, error) {
	viper := config.Load()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(viper.GetString("secret.jwt-secret")), nil
	})

	if err != nil {
		return jwt.MapClaims{}, errors.New("Token Invalid!")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return claims, errors.New("Token Expired!")
		}

		return claims, nil
	} else {
		return claims, errors.New("Token Invalid!")
	}
}
