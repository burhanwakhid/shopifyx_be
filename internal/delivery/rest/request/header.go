package request

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Header struct {
	UserId string
}

func ParseHeader(ctx *fiber.Ctx) *Header {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["user_id"].(string)

	return &Header{
		UserId: id,
	}
}
