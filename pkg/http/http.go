// By Saif Hamdan, Team Lead
// Date: 2025/1/6
package http

import (
	"github.com/goccy/go-json"
	"github.com/saifhamdan/go-apigateway-blueprint/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

const (
	StatusBadRequest          = fiber.StatusBadRequest
	StatusUnauthorized        = fiber.StatusUnauthorized
	StatusForbidden           = fiber.StatusForbidden
	StatusNotFound            = fiber.StatusNotFound
	StatusInternalServerError = fiber.StatusInternalServerError
	StatusOK                  = fiber.StatusOK
	StatusCreated             = fiber.StatusCreated
	StatusNoContent           = fiber.StatusNoContent
	StatusTooManyRequests     = fiber.StatusTooManyRequests
)

const (
	ErrBadRequest          = "Bad request"
	ErrInternalServerError = "Internal server error"
	ErrAlreadyExists       = "Already exists"
	ErrNotFound            = "Not Found"
	ErrUnauthorized        = "Unauthorized"
	ErrForbidden           = "Forbidden"
	ErrBadQueryParams      = "Invalid query params"
	ErrRequestTimeout      = "Request Timeout"
	ErrEndpointNotFound    = "The endpoint you requested doesn't exist on server"
)

type App struct {
	// fiber app instence
	*fiber.App
	// logger
	Log *logger.Logger
}

func NewApp(log *logger.Logger) *App {
	newapp := fiber.New(fiber.Config{
		JSONEncoder:             json.Marshal,
		JSONDecoder:             json.Unmarshal,
		EnableTrustedProxyCheck: true,
	})

	return &App{
		App: newapp,
		Log: log,
	}
}

type HttpResponse struct {
	// Response flag indicates whether the HTTP request was successful or not
	Success bool `json:"success"`
	// Http status Code
	Code int `json:"code"`
	// if the request were successful the data will be saved here
	Data interface{} `json:"data"`
	// Generic General Error Message defined in the system
	Error string `json:"error"`
	// More detailed error message indicates why the request was unsuccessful
	Message string `json:"message"`
}

type WSResponse struct {
	// Event Reason
	EventReason string
	// Response flag indicates whether the HTTP request was successful or not
	Success bool `json:"success"`
	// Http status Code
	Code int `json:"code"`
	// if the request were successful the data will be saved here
	Data interface{} `json:"data"`
	// Generic General Error Message defined in the system
	Error string `json:"error"`
	// More detailed error message indicates why the request was unsuccessful
	Message string `json:"message"`
}

// http 200 ok http response
func (a *App) HttpResponseOK(c *fiber.Ctx, data interface{}) error {
	return c.Status(StatusOK).JSON(
		&HttpResponse{
			Success: true,
			Code:    StatusOK,
			Data:    data,
			Error:   "",
			Message: "",
		})
}

// http 201 created http response
func (a *App) HttpResponseCreated(c *fiber.Ctx, data interface{}) error {
	return c.Status(StatusCreated).JSON(
		&HttpResponse{
			Success: true,
			Code:    StatusCreated,
			Data:    data,
			Error:   "",
			Message: "",
		})
}

// http 204 no content http response
func (a *App) HttpResponseNoContent(c *fiber.Ctx) error {
	return c.Status(StatusNoContent).JSON(
		&HttpResponse{
			Success: true,
			Code:    StatusNoContent,
			Data:    nil,
			Error:   "",
			Message: "",
		})
}

// http 400 bad request http response
func (a *App) HttpResponseBadRequest(c *fiber.Ctx, message error) error {
	a.Log.Error(message.Error())
	return c.Status(StatusBadRequest).JSON(
		&HttpResponse{
			Success: false,
			Code:    StatusBadRequest,
			Data:    nil,
			Error:   ErrBadRequest,
			Message: message.Error(),
		})
}

// http 400 bad query params http response
func (a *App) HttpResponseBadQueryParams(c *fiber.Ctx, message error) error {
	a.Log.Error(message.Error())
	return c.Status(StatusBadRequest).JSON(
		&HttpResponse{
			Success: false,
			Code:    StatusBadRequest,
			Data:    nil,
			Error:   ErrBadQueryParams,
			Message: message.Error(),
		})
}

// http 404 not found http response
func (a *App) HttpResponseNotFound(c *fiber.Ctx, message error) error {
	a.Log.Error(message.Error())
	return c.Status(StatusNotFound).JSON(
		&HttpResponse{
			Success: false,
			Code:    StatusNotFound,
			Data:    nil,
			Error:   ErrNotFound,
			Message: message.Error(),
		})
}

// http 500 internal server error response
func (a *App) HttpResponseInternalServerErrorRequest(c *fiber.Ctx, message error) error {
	a.Log.Error(message.Error())
	return c.Status(StatusInternalServerError).JSON(
		&HttpResponse{
			Success: false,
			Code:    StatusInternalServerError,
			Data:    nil,
			Error:   ErrInternalServerError,
			Message: message.Error(),
		})
}

// http 403 The client does not have access rights to the content;
// that is, it is unauthorized, so the server is refusing to give the requested resource
func (a *App) HttpResponseForbidden(c *fiber.Ctx, message error) error {
	a.Log.Error(message.Error())
	return c.Status(StatusForbidden).JSON(
		&HttpResponse{
			Success: false,
			Code:    StatusForbidden,
			Data:    nil,
			Error:   ErrForbidden,
			Message: message.Error(),
		})
}

// http 401 the client must authenticate itself to get the requested response
func (a *App) HttpResponseUnauthorized(c *fiber.Ctx, message error) error {
	a.Log.Error(message.Error())
	return c.Status(StatusUnauthorized).JSON(
		&HttpResponse{
			Success: false,
			Code:    StatusUnauthorized,
			Data:    nil,
			Error:   ErrUnauthorized,
			Message: message.Error(),
		})
}

// http 400 bad request http response
func (a *App) HttpResponseTooManyRequests(c *fiber.Ctx) error {
	a.Log.Error(fiber.ErrTooManyRequests)
	return c.Status(StatusTooManyRequests).JSON(
		&HttpResponse{
			Success: false,
			Code:    StatusBadRequest,
			Data:    nil,
			Error:   ErrBadRequest,
			Message: fiber.ErrTooManyRequests.Error(),
		})
}

// http 200 retrieve File response
func (a *App) HttpResponseFile(c *fiber.Ctx, file []byte) error {
	return c.Status(fiber.StatusOK).Send(file)
}
