package apiv1

import (
	"feed_collector/slogger"
	"feed_collector/usecase/collect_feed_usecase"

	"github.com/labstack/echo/v4"
)

type targetURL struct {
	URL string `json:"url"`
}

func CollectSingleFeed(e echo.Context) error {
	var reqString targetURL

	err := e.Bind(&reqString)
	if err != nil {
		slogger.Logger.Error("Failed to bind the request body.", "Caused by", err)
		return e.JSON(400, "Failed to bind the request body.")
	}

	err = collect_feed_usecase.CollectSingleFeedUsecase(reqString.URL)
	if err != nil {
		slogger.Logger.Error("Failed to collect the feed.", "Caused by", err)
		return e.JSON(500, "Failed to collect the feed.")
	}

	return e.JSON(200, "Successfully collected the feed.")
}
