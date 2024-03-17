package middleware

import (
	"github.com/burhanwakhid/shopifyx_backend/config"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func JwtRestAuth() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.GetJwtConfig().Secret)},
	})
}
