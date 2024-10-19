package session

import (
	cdtime "cookdroogers/pkg/time"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("ultra-secret-key")

func GetAuthenticatedUser(r *http.Request) (string, string, error) {
	authHeader := r.Header.Get("access_token")
	if authHeader == "" {
		return "", "", fmt.Errorf("missing Authorization header")
	}

	token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		return "", "", fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"].(string)
		role := claims["role"].(string)
		return email, role, nil
	} else {
		return "", "", fmt.Errorf("invalid token claims")
	}
}

func CreateToken(email, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(cdtime.Week).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (string, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return "", "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"].(string)
		role := claims["role"].(string)
		return email, role, nil
	}

	return "", "", fmt.Errorf("invalid token")
}
