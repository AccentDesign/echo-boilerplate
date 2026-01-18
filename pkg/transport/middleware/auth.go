package middleware

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v5"
)

// Authenticated middleware function to ensure the user is logged in; redirects to login if not.
func Authenticated() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			cc := c.Get("custom_context").(*CustomContext)

			setCurrentUser(c, cc)

			if c.Get("user") == nil {
				if cc.IsHTMXRequest() {
					cc.HTMXRedirect("/auth/login")
					return c.NoContent(http.StatusNoContent)
				}
				return c.Redirect(http.StatusFound, "/auth/login")
			}

			return next(c)
		}
	}
}

// setCurrentUser sets the current active user in the context.
func setCurrentUser(c *echo.Context, cc *CustomContext) {
	userIDInterface, ok := cc.Session.Values["user_id"]
	if !ok {
		return
	}

	userID, ok := userIDInterface.([16]byte)
	if !ok {
		return
	}

	ctx := c.Request().Context()
	user, err := cc.Queries.GetUserByID(ctx, pgtype.UUID{Bytes: userID, Valid: true})
	if err != nil || !user.IsActive {
		return
	}

	c.Set("user", user)
}
