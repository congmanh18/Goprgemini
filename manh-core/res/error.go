package res

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// Err { status: "ERR_500", message: "invalid request" }
type Err struct {
	Status  string   `json:"status,omitempty"`
	Message string   `json:"message,omitempty"`
	Errors  []string `json:"errors,omitempty"`
}

// Echo frameworks
func ResponseErr(c echo.Context, code int, message ...string) error {
	if len(message) == 0 {
		return c.JSON(code, Err{
			Status: fmt.Sprintf("ERR_%d", code),
		})

	}
	return c.JSON(code, Err{
		Status:  fmt.Sprintf("ERR_%d", code),
		Message: message[0],
	})
}

// Echo frameworks
func ResponseErrs(c echo.Context, code int, err error, messages ...string) error {
	var resMsg string
	if len(messages) == 0 {
		resMsg = http.StatusText(code)
	} else {
		resMsg = messages[0]
	}

	var errs = strings.Split(err.Error(), "\n")
	return c.JSON(code, Err{
		Status:  fmt.Sprintf("ERR_%d", code),
		Message: resMsg,
		Errors:  errs,
	})

}
