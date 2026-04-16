package response

import "github.com/gofiber/fiber/v3"

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func JSON(c fiber.Ctx, status int, success bool, message string, data any) error {
	return c.Status(status).JSON(Response{
		Success: success,
		Message: message,
		Data:    data,
	})
}

func Success(message string, data ...any) Response {
	var d any
	if len(data) > 0 {
		d = data[0]
	}
	return Response{
		Success: true,
		Message: message,
		Data:    d,
	}
}

func Error(message string) Response {
	return Response{
		Success: false,
		Message: message,
	}
}
