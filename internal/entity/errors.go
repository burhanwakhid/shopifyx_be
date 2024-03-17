package entity

import (
	"errors"
	"fmt"
	"net/http"
)

type CryptoError interface {
	error
	GetCode() string
	GetStatusCode() int
	GetReason() string
	GetVars() []interface{}
}

type GeneralError struct {
	Code    string
	Message string
	Reason  string
	Vars    []interface{}
}

// GeneralError function return message set when error created
func (e *GeneralError) Error() string {
	return e.Message
}

func (e *GeneralError) GetReason() string {
	return e.Reason
}

func (e *GeneralError) GetCode() string {
	return e.Code
}

func (e *GeneralError) GetVars() []interface{} {
	return e.Vars
}

// GetStatusCode function cast error interface to
func (e *GeneralError) GetStatusCode() int {
	statusCode, ok := MapCodeStatus[e.Code]
	if !ok {
		return DefaultStatus
	}

	return statusCode
}

// NewGeneralError function return error interface compatible struct with non formatted error message
func NewGeneralError(
	code string,
	message, reason string,
	vars ...interface{},
) *GeneralError {
	return &GeneralError{
		Code:    code,
		Message: message,
		Reason:  reason,
		Vars:    vars,
	}
}

func NewGeneralErrors(
	code string,
	message, reason string,
	params ...interface{},
) []*GeneralError {
	return []*GeneralError{
		NewGeneralError(
			code,
			message,
			reason,
			params...,
		),
	}
}

// NewErrorf function return error interface compatible struct with formatted error message
func NewErrorf(code string, message string, args ...interface{}) *GeneralError {
	return &GeneralError{
		Code:    code,
		Message: fmt.Sprintf(message, args...),
	}
}

// GetErrorCode function cast error interface to
func GetErrorCode(err error) *GeneralError {
	errorCode, ok := err.(*GeneralError)
	if !ok {
		return &GeneralError{}
	}

	return errorCode
}

// SetCode function
func SetCode(err error, code string) *GeneralError {
	newErr := createNewIfNotNil(err)
	newErr.Code = code

	return newErr
}

// SetReason function
func SetReason(err error, reason string) *GeneralError {
	newErr := createNewIfNotNil(err)
	newErr.Reason = reason

	return newErr
}

// SetMessage function
func SetMessage(err error, message string) *GeneralError {
	newErr := createNewIfNotNil(err)
	newErr.Message = message

	return newErr
}

func createNewIfNotNil(err error) *GeneralError {
	errorCode, ok := err.(*GeneralError)
	if ok {
		return errorCode
	}

	return &GeneralError{}
}

// WrapGoError wraps the built-in Go error with another error so that we can use errors.Is and errors.As methods.
func WrapGoError(err1, err2 error) error {
	return fmt.Errorf("%w: %s", err1, err2.Error())
}

var DefaultStatus = http.StatusInternalServerError

var MapCodeStatus = map[string]int{
	CodeInvalidParameter: http.StatusBadRequest,
	CodeRedisError:       http.StatusInternalServerError,
	CodePostgresError:    http.StatusInternalServerError,
	CodeExternalAPIError: http.StatusInternalServerError,
}

var (
	// General
	CodeInvalidParameter           string = "00001"
	CodeRedisError                 string = "00002"
	CodePostgresError              string = "00003"
	CodeExternalAPIError           string = "00004"
	CodeAlreadyExists              string = "00006"
	CodePINIsRequired              string = "00007"
	CodePINTokenIsInvalidOrExpired string = "00008"
	CodeInsufficientGopayBalance   string = "00009"
	CodeInsufficientAssetBalance   string = "00010"
	CodeAmountBelowMinLimit        string = "00011"
	CodeAmountAboveMaxLimit        string = "00012"
	CodeKycNotVerified             string = "00013"
	InvalidPassword                string = "00014"

	CodeInvalidRequest string = "31005"
	CodeUnauthorized   string = "00401"
	CodeNotFound       string = "00404"
	CodeTimeout        string = "00408"
	CodeInternalServer string = "00500"
)

// GeneralError Code Template which contains error code and the message
var (
	ErrInvalidParameter  *GeneralError = NewGeneralError(CodeInvalidParameter, "invalid parameter", "")
	ErrRedisError        *GeneralError = NewGeneralError(CodeRedisError, "internal server error", "")
	ErrPostgres          *GeneralError = NewGeneralError(CodePostgresError, "internal server error", "")
	ErrExternalAPI       *GeneralError = NewGeneralError(CodeExternalAPIError, "internal server error", "")
	ErrDataNotFound      *GeneralError = NewGeneralError(CodeNotFound, "data not found", "")
	ErrInvalidPassword   *GeneralError = NewGeneralError(InvalidPassword, "Invalid Password", "")
	ErrDataAlreadyExists *GeneralError = NewGeneralError(CodeAlreadyExists, "data already exists", "")
)

// Golang errors.
var (
	ErrHTTPRequest       = errors.New("http request error")
	ErrJSONMarshal       = errors.New("json marshal error")
	ErrJSONUnmarshal     = errors.New("json unmarshal error")
	ErrExternalAPICall   = errors.New("external api error")
	ErrRequestValidation = errors.New("request validation error")
	ErrInvalidVendorID   = errors.New("invalid vendor ID")
	ErrCreateDBRecord    = errors.New("error creating db record")
	ErrDBTx              = errors.New("db transaction error")
	ErrUpdateDBRecord    = errors.New("error updating db record")
	ErrGetDBRecord       = errors.New("error getting db record")
)

var (
	ErrHTTPRequestTimeout = NewGeneralError(CodeTimeout, "request timeout to external party", "")
	ErrHTTPInternalServer = NewGeneralError(CodeInternalServer, "request internal server error to external party", "")
	ErrHTTPNotFound       = NewGeneralError(CodeNotFound, "request not found to external party", "")
	ErrHTTPUnauthorized   = NewGeneralError(CodeUnauthorized, "request unauthorized to external party", "")
	ErrHTTPBadRequest     = NewGeneralError(CodeInvalidParameter, "request bad request to external party", "")
)

var (
	MapHTTPStatusToError = map[int]error{
		http.StatusRequestTimeout:      ErrHTTPRequestTimeout,
		http.StatusInternalServerError: ErrHTTPInternalServer,
		http.StatusNotFound:            ErrHTTPNotFound,
		http.StatusUnauthorized:        ErrHTTPUnauthorized,
		http.StatusBadRequest:          ErrHTTPBadRequest,
	}
)

func HTTPStatusToError(status int) error {
	e, ok := MapHTTPStatusToError[status]
	if !ok {
		return ErrHTTPInternalServer
	}

	return e
}
