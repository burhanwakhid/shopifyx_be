package token

import (
	"time"

	"github.com/burhanwakhid/shopifyx_backend/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt(username, userId string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"user_id":  userId,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := config.GetJwtConfig().Secret

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return t, nil
}
