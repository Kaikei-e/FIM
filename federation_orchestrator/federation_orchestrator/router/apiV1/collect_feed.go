package apiv1

import (
	"context"
	"federation_orchestrator/slogger"
	"fmt"

	"github.com/labstack/echo/v4"
)

type targetURL struct {
	URL string `json:"url"`
}

func CollectSingleFeed(e echo.Context, ctx context.Context) error {
	var reqString targetURL

	err := e.Bind(&reqString)
	if err != nil {
		slogger.Logger.Error("Failed to bind the request body.", "Caused by", err)
		return e.JSON(400, "Failed to bind the request body.")
	}

	fmt.Println(reqString)

	return e.JSON(200, "Successfully collected the feed.")
}
