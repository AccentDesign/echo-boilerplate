package middleware

import (
	"echo.go.dev/pkg/config"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

// CSRF returns an Echo middleware function for csrf.
func CSRF(cfg *config.Config) echo.MiddlewareFunc {
	csrfConfig := middleware.CSRFConfig{
		TrustedOrigins: cfg.Session.TrustedOrigins,
		CookiePath:     cfg.Session.Path,
		CookieDomain:   cfg.Session.Domain,
		CookieSecure:   cfg.Session.Secure,
		CookieHTTPOnly: cfg.Session.HttpOnly,
		CookieSameSite: cfg.Session.SameSite,
	}

	return middleware.CSRFWithConfig(csrfConfig)
}
