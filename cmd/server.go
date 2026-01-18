package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"echo.go.dev/pkg/config"
	"echo.go.dev/pkg/domain"
	apperrors "echo.go.dev/pkg/errors"
	"echo.go.dev/pkg/static"
	"echo.go.dev/pkg/storage/db"
	"echo.go.dev/pkg/transport/middleware"
	"echo.go.dev/pkg/transport/validate"
	"github.com/labstack/echo/v5"
	echomiddleware "github.com/labstack/echo/v5/middleware"
	"github.com/spf13/cobra"
)

var cmdServer = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

func runServer() {
	cfg := config.GetConfig()
	ctx := context.Background()
	dbPool, err := db.GetPool(ctx)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer dbPool.Close()

	engine := echo.New()
	engine.Validator = validate.New()
	engine.HTTPErrorHandler = apperrors.HttpErrorHandler

	engine.Use(
		echomiddleware.Recover(),
		middleware.Logger(),
		middleware.Secure(cfg),
		middleware.CSRF(cfg),
		middleware.CORS(cfg),
		middleware.RateLimit(),
		middleware.Gzip(),
		middleware.Session(cfg),
		middleware.Context(dbPool),
	)

	static.Router(engine)
	domain.Router(engine)

	server := http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Server.Port),
		ReadHeaderTimeout: 10 * time.Second,
		Handler:           engine,
	}

	log.Printf("Starting server on port %d", cfg.Server.Port)
	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}
}
