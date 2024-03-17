package response

import (
	"errors"
	"net/http"

	"github.com/burhanwakhid/shopifyx_backend/internal"
	"github.com/burhanwakhid/shopifyx_backend/internal/entity"
	"github.com/gofiber/fiber/v2"
)

func SendResponseErrorInternalServer(
	c *fiber.Ctx,
	t internal.ErrorTranslator,
	locale string,
	data interface{},
	err error,
) error {
	return SendResponseParseErrorCode(c, t, locale, data, err, "500", nil)
}

func SendResponseCustomCodeError(
	c *fiber.Ctx,
	t internal.ErrorTranslator,
	locale string,
	code string,
	data interface{},
	err error,
) error {

	vars := []interface{}{}

	return SendResponseParseErrorCode(c, t, locale, data, err, code, vars)
}

func SendResponseErrorBadParameter(
	c *fiber.Ctx,
	t internal.ErrorTranslator,
	locale string,
	data interface{},
	err error,
) error {
	code := "500"
	return SendResponseParseErrorCode(c, t, locale, data, err, code, nil)
}

func SendResponseErrorForbidden(
	c *fiber.Ctx,
	t internal.ErrorTranslator,
	locale string,
	data interface{},
	err error,
) error {
	code := "401"
	return SendResponseParseErrorCode(c, t, locale, data, err, code, nil)
}

func SendResponseErrorBadRequest(
	c *fiber.Ctx,
	t internal.ErrorTranslator,
	locale string,
	data interface{},
	err error,
) error {
	code := "400"
	return SendResponseParseErrorCode(c, t, locale, data, err, code, nil)
}

func SendResponseParseErrorCode(
	c *fiber.Ctx,
	t internal.ErrorTranslator,
	locale string,
	data interface{},
	err error,
	code string,
	vars []interface{},
) error {
	template := t.Translate(locale, code, vars...)

	c.Status(template.HTTPStatusCode)

	return c.JSON(

		entity.NewRestResponse(
			false,
			data,
			nil,
			NewError(
				err,
				code,
				template.MessageTitle,
				template.Message,
				template.Severity,
			),
		),
	)
}

func SendResponse(
	c *fiber.Ctx,
	t internal.ErrorTranslator,
	locale string,
	data interface{},
	meta interface{},
	errs ...*entity.GeneralError,
) error {
	var (
		success    = true
		statusCode = http.StatusOK
		errResp    = make([]entity.RestError, 0, len(errs))
	)

	if len(errs) > 0 {
		success = false
		statusCode = t.Translate(locale, errs[0].Code).HTTPStatusCode

		for i := range errs {
			err := t.Translate(locale, errs[i].Code, errs[i].Vars...)

			errResp = append(
				errResp,
				NewError(
					errors.New(errs[i].Reason),
					err.Code,
					err.MessageTitle,
					err.Message,
					err.Severity,
				),
			)
		}
	}
	c.Status(statusCode)

	err := c.JSON(entity.NewRestResponse(success, data, meta, errResp...))
	if err != nil {

		return SendResponseErrorInternalServer(c, t, locale, nil, err)
	}
	return nil
}

func SendResponseCreated(
	c *fiber.Ctx,
	t internal.ErrorTranslator,
	locale string,
	data interface{},
	meta interface{},
	errs ...*entity.GeneralError,
) error {
	var (
		success    = true
		statusCode = http.StatusCreated
		errResp    = make([]entity.RestError, 0, len(errs))
	)

	if len(errs) > 0 {
		success = false
		statusCode = t.Translate(locale, errs[0].Code).HTTPStatusCode

		for i := range errs {
			err := t.Translate(locale, errs[i].Code, errs[i].Vars...)

			errResp = append(
				errResp,
				NewError(
					errors.New(errs[i].Reason),
					err.Code,
					err.MessageTitle,
					err.Message,
					err.Severity,
				),
			)
		}
	}
	c.Status(statusCode)

	err := c.JSON(entity.NewRestResponse(success, data, meta, errResp...))
	if err != nil {

		return SendResponseErrorInternalServer(c, t, locale, nil, err)
	}
	return nil
}
