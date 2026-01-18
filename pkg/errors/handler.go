package errors

import (
	"errors"
	"net/http"

	"echo.go.dev/pkg/transport/middleware"
	"echo.go.dev/pkg/ui/pages"
	"github.com/labstack/echo/v5"
)

func HttpErrorHandler(c *echo.Context, err error) {
	if resp, uErr := echo.UnwrapResponse(c.Response()); uErr == nil {
		if resp.Committed {
			return
		}
	}

	code := http.StatusInternalServerError
	var sc echo.HTTPStatusCoder
	if errors.As(err, &sc) {
		if tmp := sc.StatusCode(); tmp != 0 {
			code = tmp
		}
	}

	message := "An unexpected error occurred on the server."
	switch code {
	case http.StatusNotFound:
		message = "The requested URL was not found on this server."
	}

	cc := c.Get("custom_context").(*middleware.CustomContext)
	if err := cc.RenderComponent(code, pages.Error(pages.ErrorProps{
		Code:    code,
		Title:   http.StatusText(code),
		Message: message,
	})); err != nil {
		c.Logger().Error(err.Error())
	}
}
