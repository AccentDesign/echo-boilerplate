package middleware

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

// Gzip returns an Echo middleware function for gzip.
func Gzip() echo.MiddlewareFunc {
	gzipConfig := middleware.GzipConfig{
		Level:     0,
		MinLength: 0,
		Skipper: func(c *echo.Context) bool {
			return c.Path() != "/static*"
		},
	}

	return middleware.GzipWithConfig(gzipConfig)
}
