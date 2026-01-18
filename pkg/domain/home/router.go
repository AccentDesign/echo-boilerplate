package home

import (
	"net/http"

	"echo.go.dev/pkg/transport/middleware"
	"echo.go.dev/pkg/ui/pages"
	"github.com/labstack/echo/v5"
)

// Router create a new Router.
func Router(e *echo.Echo) {
	auth := middleware.Authenticated()
	g := e.Group("", auth)
	{
		g.GET("", index)
	}
}

// index root page endpoint.
func index(c *echo.Context) error {
	cc := c.Get("custom_context").(*middleware.CustomContext)
	page := pages.Page{Title: "Home"}
	return cc.RenderComponent(http.StatusOK, pages.Home(page))
}
