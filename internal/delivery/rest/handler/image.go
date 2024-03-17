package handler

import (
	"github.com/burhanwakhid/shopifyx_backend/internal"
	"github.com/burhanwakhid/shopifyx_backend/internal/delivery/rest/response"
	"github.com/burhanwakhid/shopifyx_backend/internal/service"
	validates "github.com/burhanwakhid/shopifyx_backend/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type ImageHandler struct {
	service         service.ImageService
	errorTranslator internal.ErrorTranslator
	validate        validates.JSONValidator
}

func NewImageHandler(
	service service.ImageService,
	errorTranslator internal.ErrorTranslator,
	validate validates.JSONValidator,
) *ImageHandler {
	return &ImageHandler{
		service:         service,
		errorTranslator: errorTranslator,
		validate:        validate,
	}
}

func (handler *ImageHandler) Mount(group fiber.Router) {
	group.Post("", handler.UploadImage)
}

func (handler *ImageHandler) UploadImage(c *fiber.Ctx) error {

	file, err := c.FormFile("file")

	if err != nil {
		return response.SendResponseErrorBadRequest(c, handler.errorTranslator, "en", nil, err)
	}

	url, err := handler.service.UploadImage(file)

	if err != nil {
		return response.SendResponseErrorBadRequest(c, handler.errorTranslator, "en", nil, err)
	}

	return response.SendResponse(c, handler.errorTranslator, "en", response.ImageResponse{ImageUrl: url}, nil)
}
