package middleware

import (
	"context"
	"log/slog"
	"os"

	"echo.go.dev/pkg/storage/db/dbx"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func Logger() echo.MiddlewareFunc {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:       true,
		LogStatus:       true,
		LogProtocol:     true,
		LogHost:         true,
		LogURI:          true,
		LogURIPath:      true,
		LogRoutePath:    true,
		LogRemoteIP:     true,
		LogReferer:      true,
		LogLatency:      true,
		LogResponseSize: true,
		LogUserAgent:    true,
		HandleError:     true,
		LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
			attrs := []slog.Attr{
				slog.String("method", v.Method),
				slog.Int("status", v.Status),
				slog.String("protocol", v.Protocol),
				slog.String("host", v.Host),
				slog.String("uri", v.URI),
				slog.String("uri_path", v.URIPath),
				slog.String("route_path", v.RoutePath),
				slog.String("remote_ip", v.RemoteIP),
				slog.String("referer", v.Referer),
				slog.Duration("latency", v.Latency),
				slog.Int64("response_size", v.ResponseSize),
				slog.String("user_agent", v.UserAgent),
			}
			if user, ok := c.Get("user").(dbx.AuthUser); ok {
				attrs = append(attrs, slog.String("user", user.Email))
			}
			logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST", attrs...)
			return nil
		},
		Skipper: func(c *echo.Context) bool {
			return c.Path() == "/static*"
		},
	})
}
