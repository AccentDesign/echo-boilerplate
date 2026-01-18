package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"strings"

	"echo.go.dev/pkg/config"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

// Secure returns an Echo middleware function for cors.
func Secure(cfg *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			nonce := generateNonce()
			policy := strings.ReplaceAll(cfg.Security.CSP(), "nonce-", "nonce-"+nonce)

			secureConfig := middleware.SecureConfig{
				XSSProtection:         cfg.Security.XSSProtection,
				ContentTypeNosniff:    cfg.Security.ContentTypeNosniff,
				XFrameOptions:         cfg.Security.XFrameOptions,
				HSTSMaxAge:            cfg.Security.HSTSMaxAge,
				ContentSecurityPolicy: policy,
				ReferrerPolicy:        cfg.Security.ReferrerPolicy,
				Skipper: func(c *echo.Context) bool {
					return c.Path() != "/static*"
				},
			}

			ctx := templ.WithNonce(c.Request().Context(), nonce)
			c.SetRequest(c.Request().WithContext(ctx))

			return middleware.SecureWithConfig(secureConfig)(next)(c)
		}
	}
}

// generateNonce generates a random base64 nonce.
func generateNonce() string {
	nonce := make([]byte, 16)
	_, err := rand.Read(nonce)
	if err != nil {
		panic("failed to generate nonce: " + err.Error())
	}
	return base64.StdEncoding.EncodeToString(nonce)
}
