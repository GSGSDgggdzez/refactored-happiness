package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("your-secret-key-here") // In production, use environment variable

type Claims struct {
	UserID             int64  `json:"user_id"`
	Email              string `json:"email"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Country            string `json:"Country"`
	Phone              string `json:"Phone"`
	Verification_token string `json:"Verification_token"`
	jwt.RegisteredClaims
}

func GenerateToken(userID int64, email, firstName string, lastName string, country string, phone string, verification_token string) (string, error) {
	claims := Claims{
		UserID:             userID,
		Email:              email,
		FirstName:          firstName,
		LastName:           lastName,
		Country:            country,
		Phone:              phone,
		Verification_token: verification_token,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

func ExtractTokenFromHeader(header string) (*Claims, *fiber.Map) {
	if header == "" {
		return nil, &fiber.Map{
			"status": fiber.StatusUnauthorized,
			"error":  "Authorization header is required",
		}
	}

	parts := strings.Split(header, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, &fiber.Map{
			"status": fiber.StatusUnauthorized,
			"error":  "Invalid authorization header format",
		}
	}

	tokenString := parts[1]

	claims, err := ValidateToken(tokenString)
	if err != nil {
		return nil, &fiber.Map{
			"status": fiber.StatusUnauthorized,
			"error":  fmt.Sprintf("Invalid token: %v", err),
		}
	}

	return claims, nil
}
