package res

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// OK { status: "OK_200", data: {} }
type OK struct {
	Status  string      `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseOK(c echo.Context, code int, message string, data interface{}) error {
	return c.JSON(code, OK{
		Status:  fmt.Sprintf("OK_%d", code),
		Message: message,
		Data:    data,
	})
}

func SimpleResponseOK(c echo.Context, code int, data interface{}) error {
	return c.JSON(code, OK{
		Status:  fmt.Sprintf("OK_%d", code),
		Message: "OK",
		Data:    data,
	})
}
