package session

import (
	cdtime "cookdroogers/pkg/time"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("ultra-secret-key")

func GetAuthenticatedUser(r *http.Request) (uint64, string, string, error) {
	authHeader := r.Header.Get("access_token")
	if authHeader == "" {
		return 0, "", "", fmt.Errorf("missing Authorization header")
	}

	token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		return 0, "", "", fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		idStr := claims["id"].(string)
		idInt, err := strconv.Atoi(idStr)
		if err != nil {
			return 0, "", "", fmt.Errorf("invalid token")
		}

		id := uint64(idInt)
		email := claims["email"].(string)
		role := claims["role"].(string)
		return id, email, role, nil
	}
	return 0, "", "", fmt.Errorf("invalid token claims")
}

func CreateToken(id uint64, email, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    strconv.FormatUint(id, 10),
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

func VerifyToken(tokenString string) (uint64, string, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return 0, "", "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		idStr := claims["id"].(string)
		idInt, err := strconv.Atoi(idStr)
		if err != nil {
			return 0, "", "", fmt.Errorf("invalid token")
		}

		id := uint64(idInt)
		email := claims["email"].(string)
		role := claims["role"].(string)
		return id, email, role, nil
	}

	return 0, "", "", fmt.Errorf("invalid token")
}
