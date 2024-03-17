package handler

import (
	"github.com/burhanwakhid/shopifyx_backend/internal"
	"github.com/burhanwakhid/shopifyx_backend/internal/delivery/rest/request"
	"github.com/burhanwakhid/shopifyx_backend/internal/delivery/rest/response"
	"github.com/burhanwakhid/shopifyx_backend/internal/service"
	validates "github.com/burhanwakhid/shopifyx_backend/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	service         service.ProductService
	errorTranslator internal.ErrorTranslator
	validate        validates.JSONValidator
}

func NewProductHandler(
	service service.ProductService,
	errorTranslator internal.ErrorTranslator,
	validate validates.JSONValidator,
) *ProductHandler {
	return &ProductHandler{
		service:         service,
		errorTranslator: errorTranslator,
		validate:        validate,
	}
}

func (handler *ProductHandler) Mount(group fiber.Router) {
	group.Post("", handler.CreateProduct)
	// group.Get("/account", handler.GetBankUser)
}

func (handler *ProductHandler) CreateProduct(c *fiber.Ctx) error {

	var req request.Product
	header := request.ParseHeader(c)

	if err := c.BodyParser(&req); err != nil {
		return response.SendResponseErrorBadRequest(c, handler.errorTranslator, "en", nil, err)
	}

	if err := handler.validate.Validate(c.UserContext(), req); err != nil {

		return response.SendResponseErrorBadRequest(c, handler.errorTranslator, "en", nil, err)
	}

	err := handler.service.CreateProduct(c.UserContext(), req, header.UserId)

	if err != nil {
		return response.SendResponseErrorBadRequest(c, handler.errorTranslator, "en", nil, err)
	}

	return response.SendResponse(c, handler.errorTranslator, "en", true, nil)

}
