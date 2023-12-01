package middleware

import (
	"fmt"
	"fww-wrapper/internal/repository"
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type Middleware struct {
	Repository repository.Repository
}

func (m *Middleware) ValidateAPIKey(ctx *fiber.Ctx) error {
	apiKey := ctx.Get("X-API-Key")
	partnerID := ctx.Get("X-Partner-ID")
	if apiKey == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// validate partner id
	partner, err := m.Repository.FindPartnerByID(partnerID)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// validate api key
	if partner.ApiKey != apiKey {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return ctx.Next()
}

func (m *Middleware) ValidateJWTUser(ctx *fiber.Ctx) error {
	jwtToken := ctx.Get("Authorization")
	if jwtToken == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// validate jwt token
	userID, err := decodeToken(jwtToken)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// String to Int64
	userIDInt64, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		log.Fatal(err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// validate user id
	_, err = m.Repository.FindUserByID(userIDInt64)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return ctx.Next()
}

type CustomClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func decodeToken(jwtToken string) (string, error) {
	// Decode Token JWT
	token, err := jwt.ParseWithClaims(jwtToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		if claims.StandardClaims.ExpiresAt < time.Now().Unix() {
			return "", err
		}
		return claims.UserID, nil
	} else {
		return "", err
	}

}
