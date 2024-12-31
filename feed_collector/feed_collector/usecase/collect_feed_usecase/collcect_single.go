package collect_feed_usecase

import (
	"context"
	"database/sql"
	"feed_collector/client"
	"feed_collector/domain/rss_feed"
	"feed_collector/driver/db_driver"
	"feed_collector/protector"
	"feed_collector/slogger"
)

func CollectSingleFeedUsecase(reqBody string, db *sql.DB, ctx context.Context) error {
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

	insertingFeed := rss_feed.Feed{
		URL:           feed.Link,
		Title:         feed.Title,
		Description:   feed.Description,
		Author:        feed.Author.Name,
		PublishedDate: *feed.PublishedParsed,
	}

	err = db_driver.Register(insertingFeed, db, ctx)
	if err != nil {
		slogger.Logger.Error("Failed to register the feed.", "Caused by", err)
		return err
	}

	return nil
}
