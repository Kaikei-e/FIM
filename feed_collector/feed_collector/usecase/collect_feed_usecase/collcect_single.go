package collect_feed_usecase

import (
	"feed_collector/client"
	"feed_collector/protector"
	"feed_collector/slogger"
	"fmt"
)

func CollectSingleFeedUsecase(reqBody string) error {
	cl := protector.InitClient()

	targetURL, err := protector.AntiURLAttack(reqBody, cl)
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
