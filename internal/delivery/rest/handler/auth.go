package handler

import (
	"errors"
	"fmt"

	"github.com/burhanwakhid/shopifyx_backend/internal"
	"github.com/burhanwakhid/shopifyx_backend/internal/delivery/rest/request"
	"github.com/burhanwakhid/shopifyx_backend/internal/delivery/rest/response"
	"github.com/burhanwakhid/shopifyx_backend/internal/entity"
	"github.com/burhanwakhid/shopifyx_backend/internal/service"
	validates "github.com/burhanwakhid/shopifyx_backend/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service         service.AuthService
	errorTranslator internal.ErrorTranslator
	validate        validates.JSONValidator
}

func NewAuthHandler(
	service service.AuthService,
	errorTranslator internal.ErrorTranslator,
	validate validates.JSONValidator,
) *AuthHandler {
	return &AuthHandler{
		service:         service,
		errorTranslator: errorTranslator,
		validate:        validate,
	}
}

func (handler *AuthHandler) Mount(group fiber.Router) {
	group.Post("/register", handler.RegisterUser)
	group.Post("/login", handler.Loginuser)
}

func (handler *AuthHandler) RegisterUser(c *fiber.Ctx) error {
	var req request.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return response.SendResponseErrorBadRequest(c, handler.errorTranslator, "en", nil, err)
	}

	if err := handler.validate.Validate(c.UserContext(), req); err != nil {

		return response.SendResponseErrorInternalServer(c, handler.errorTranslator, "en", nil, err)
	}

	user, err := handler.service.RegisterUser(c.UserContext(), req)

	if err != nil {
		fmt.Printf("error validate register: %s", err)
		if errors.Is(err, entity.ErrDataAlreadyExists) {
			return response.SendResponseCustomCodeError(c, handler.errorTranslator, "en", "409", nil, err)
		}
		return response.SendResponseErrorBadRequest(c, handler.errorTranslator, "en", nil, err)
	}

	return response.SendResponseCreated(c, handler.errorTranslator, "en", handler.loginTransform(user), nil)
}

func (handler *AuthHandler) Loginuser(c *fiber.Ctx) error {

	var request request.LoginRequest

	if err := c.BodyParser(&request); err != nil {
		return response.SendResponseErrorBadRequest(c, handler.errorTranslator, "en", nil, err)
	}

	if err := handler.validate.Validate(c.UserContext(), request); err != nil {
		fmt.Printf("error validate login: %s", err)
		if errors.Is(err, entity.ErrDataNotFound) {
			return response.SendResponseCustomCodeError(c, handler.errorTranslator, "en", "404", nil, err)
		}
		return response.SendResponseErrorBadRequest(c, handler.errorTranslator, "en", nil, err)
	}

	user, err := handler.service.LoginUser(c.UserContext(), request.Username, request.Password)

	if err != nil {
		return response.SendResponseErrorBadRequest(c, handler.errorTranslator, "en", nil, err)
	}

	return response.SendResponse(c, handler.errorTranslator, "en", handler.loginTransform(user), nil)

}

func (handler *AuthHandler) loginTransform(user *entity.LoginData) response.LoginResponse {

	usr := response.LoginResponse{
		Username: user.Username,
		Token:    user.Token,
		Name:     user.Name,
	}

	return usr
}
