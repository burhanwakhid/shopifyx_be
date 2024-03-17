package response

import (
	"fmt"

	"github.com/burhanwakhid/shopifyx_backend/internal/entity"
)

// var (
// 	// GopayErrorCodeToHTTPStatusCode TODO: Temporary. Move this inside the message translation.
// 	GopayErrorCodeToHTTPStatusCode = map[gopay.GoPayError]int{

// 		// Crypto generic errors.
// 		gopay.GoPayError_CRYPTO_INTERNAL_SERVER_ERROR: http.StatusInternalServerError,
// 		gopay.GoPayError_CRYPTO_INVALID_PARAMETER:     http.StatusBadRequest,
// 		gopay.GoPayError_CRYPTO_INVALID_TOKEN:         http.StatusBadRequest,
// 		gopay.GoPayError_CRYPTO_EXTERNAL_API_ERROR:    http.StatusInternalServerError,
// 		gopay.GoPayError_CRYPTO_NOT_FOUND_ERROR:       http.StatusBadRequest,
// 		gopay.GoPayError_CRYPTO_INVALID_REQUEST:       http.StatusBadRequest,
// 		gopay.GoPayError_CRYPTO_INSUFFICIENT_BALANCE:  http.StatusBadRequest,
// 		// Crypto specific errors.
// 		gopay.GoPayError_CRYPTO_INVALID_DESTINATION_ADDRESS:        http.StatusBadRequest,
// 		gopay.GoPayError_CRYPTO_INVALID_SENDER_PHYSICAL_ADDRESS:    http.StatusBadRequest,
// 		gopay.GoPayError_CRYPTO_INVALID_RECIPIENT_PHYSICAL_ADDRESS: http.StatusBadRequest,
// 		gopay.GoPayError_CRYPTO_INVALID_SENDER_NAME:                http.StatusBadRequest,
// 		gopay.GoPayError_CRYPTO_INVALID_RECIPIENT_NAME:             http.StatusBadRequest,
// 		gopay.GoPayError_CRYPTO_INVALID_WITHDRAWAL_AMOUNT:          http.StatusBadRequest,
// 		gopay.GoPayError_CRYPTO_WITHDRAWAL_AMOUNT_LIMIT_EXCEEDED:   http.StatusBadRequest,
// 	}
// )

// NewError TODO: Add message translation.
func NewError(
	err error,
	code string,
	messageTitle, message, severity string,
) entity.RestError {
	return entity.RestError{
		Code:            fmt.Sprintf("%s%s", "Gemini-", code),
		Message:         message,
		MessageTitle:    messageTitle,
		MessageSeverity: severity,
		Entity:          "gemini-backend",
		Cause:           err.Error(),
	}
}

// func GetHTTPStatusCode(errorCode gopay.GoPayError) int {
// 	if statusCode, exists := GopayErrorCodeToHTTPStatusCode[errorCode]; exists {
// 		return statusCode
// 	}
// 	return http.StatusInternalServerError
// }
