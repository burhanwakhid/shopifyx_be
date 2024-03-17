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

type BankHandler struct {
	service         service.BankService
	errorTranslator internal.ErrorTranslator
	validate        validates.JSONValidator
}

func NewBankHandler(
	service service.BankService,
	errorTranslator internal.ErrorTranslator,
	validate validates.JSONValidator,
) *BankHandler {
	return &BankHandler{
		service:         service,
		errorTranslator: errorTranslator,
		validate:        validate,
	}
}

func (handler *BankHandler) Mount(group fiber.Router) {
	group.Post("/account", handler.CreateBank)
	group.Get("/account", handler.GetBankUser)
	group.Patch("/account/:bankAccountId", handler.UpdateBank)
	group.Delete("/account/:bankAccountId", handler.DeleteBank)
}

func (handler *BankHandler) CreateBank(c *fiber.Ctx) error {

	var req request.BankRequest
	header := request.ParseHeader(c)

	if err := c.BodyParser(&req); err != nil {
		return response.SendResponseErrorBadRequest(c, handler.errorTranslator, "en", nil, err)
	}

	if err := handler.validate.Validate(c.UserContext(), req); err != nil {

		return response.SendResponseErrorBadRequest(c, handler.errorTranslator, "en", nil, err)
	}

	err := handler.service.CreateBank(c.UserContext(), req, header.UserId)

	if err != nil {
		return response.SendResponseErrorBadRequest(c, handler.errorTranslator, "en", nil, err)
	}

	return response.SendResponse(c, handler.errorTranslator, "en", true, nil)

}

func (handler *BankHandler) UpdateBank(c *fiber.Ctx) error {

	var req request.BankRequest

	bankAccountId := c.Params("bankAccountId")

	if err := c.BodyParser(&req); err != nil {

		return response.SendResponseErrorBadRequest(c, handler.errorTranslator, "en", nil, err)
	}

	if err := handler.validate.Validate(c.UserContext(), req); err != nil {

		return response.SendResponseErrorBadRequest(c, handler.errorTranslator, "en", nil, err)
	}

	bank, err := handler.service.UpdateBank(c.UserContext(), req, bankAccountId)

	if err != nil {

		if errors.Is(err, entity.ErrDataNotFound) {
			fmt.Printf("error validate update: %s", err)
			return response.SendResponseCustomCodeError(c, handler.errorTranslator, "en", "404", nil, err)
		}

		return response.SendResponseErrorBadRequest(c, handler.errorTranslator, "en", nil, err)
	}

	return response.SendResponse(c, handler.errorTranslator, "en", handler.conversationBankTransform(bank), nil)

}

func (handler *BankHandler) GetBankUser(c *fiber.Ctx) error {

	header := request.ParseHeader(c)

	banks, err := handler.service.GetBank(c.UserContext(), header.UserId)

	if err != nil {
		return response.SendResponseErrorBadRequest(c, handler.errorTranslator, "en", nil, err)
	}

	resBank := handler.conversationTransform(banks)

	return response.SendResponse(c, handler.errorTranslator, "en", resBank, nil)

}

func (handler *BankHandler) DeleteBank(c *fiber.Ctx) error {

	bankAccountId := c.Params("bankAccountId")

	err := handler.service.DeleteBank(c.UserContext(), bankAccountId)

	if err != nil {

		if errors.Is(err, entity.ErrDataNotFound) {
			fmt.Printf("error validate update: %s", err)
			return response.SendResponseCustomCodeError(c, handler.errorTranslator, "en", "404", nil, err)
		}

		return response.SendResponseErrorBadRequest(c, handler.errorTranslator, "en", nil, err)
	}

	return response.SendResponse(c, handler.errorTranslator, "en", nil, nil)

}

func (handler *BankHandler) conversationTransform(msg []*entity.Bank) []response.BankResponse {

	responseMessage := make([]response.BankResponse, 0, len(msg))

	for _, v := range msg {
		responseMessage = append(responseMessage, response.BankResponse{
			Id:            v.Id,
			AccountName:   v.AccountName,
			Name:          v.Name,
			AccountNumber: v.AccountNumber,
		})
	}

	return responseMessage
}
func (handler *BankHandler) conversationBankTransform(msg *entity.Bank) response.BankResponse {

	responseMessage := response.BankResponse{
		Id:            msg.Id,
		AccountName:   msg.AccountName,
		Name:          msg.Name,
		AccountNumber: msg.AccountNumber,
	}

	return responseMessage
}
