package middleware

import (
	"echo.go.dev/pkg/config"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v5"
)

// Session returns an Echo middleware function for the session.
func Session(cfg *config.Config) echo.MiddlewareFunc {
	sessionStore := sessions.NewCookieStore(cfg.Session.KeyBytes(), cfg.Session.EncKeyBytes())

	sessionStore.Options = &sessions.Options{
		Path:     cfg.Session.Path,
		Domain:   cfg.Session.Domain,
		MaxAge:   cfg.Session.MaxAge,
		Secure:   cfg.Session.Secure,
		HttpOnly: cfg.Session.HttpOnly,
		SameSite: cfg.Session.SameSite,
	}

	return session.Middleware(sessionStore)
}
