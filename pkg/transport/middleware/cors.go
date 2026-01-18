package middleware

import (
	"echo.go.dev/pkg/config"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

// CORS returns an Echo middleware function for cors.
func CORS(cfg *config.Config) echo.MiddlewareFunc {
	corsConfig := middleware.CORSConfig{
		AllowOrigins: cfg.Security.AllowOrigins,
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
		},
	}

	return middleware.CORSWithConfig(corsConfig)
}
