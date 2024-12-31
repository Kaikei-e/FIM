package router

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"

	apiv1 "feed_collector/router/apiV1"
	"feed_collector/slogger"
	logger "feed_collector/slogger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router(db *sql.DB, ctx context.Context) {
	middleware.SecureWithConfig(middleware.DefaultSecureConfig)

	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.Logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
				)
			} else {
				logger.Logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))

	apiV1Group := e.Group("/api/v1")
	{
		apiV1Group.GET("/healthCheck", healthCheck)
		apiV1Group.POST("/feeds/collect/single", func(c echo.Context) error {
			return apiv1.CollectSingleFeed(c, db, ctx)
		})
	}

	slogger.Logger.Info("Starting server on port 8000")
	e.Logger.Fatal(e.Start(":8000"))
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
