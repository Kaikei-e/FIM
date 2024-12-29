package apiv1

import (
	"feed_collector/client"
	"feed_collector/protector"
	"feed_collector/slogger"
	"fmt"

	"github.com/labstack/echo/v4"
)

func CollectSingleFeed(e echo.Context) error {
	cl := protector.InitClient()

	targetURL, err := protector.AntiURLAttack("https://lorem-rss.herokuapp.com/feed", cl)
	if err != nil {
		slogger.Logger.Error("Failed to parse the URL.", "Caused by", err)
		return err
	}

	feed, err := client.CollectRSSFeed(*targetURL, cl)
	if err != nil {
		slogger.Logger.Error("Failed to collect the feed.", "Caused by", err)
		return err
	}

	fmt.Println(feed.Link)

	return nil
}
