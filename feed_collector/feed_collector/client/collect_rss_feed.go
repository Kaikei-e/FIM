package client

import (
	"feed_collector/slogger"
	"net/url"

	"github.com/doyensec/safeurl"
	"github.com/mmcdole/gofeed"
)

func CollectRSSFeed(rssLink url.URL, cl *safeurl.WrappedClient) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(rssLink.String())
	if err != nil {
		slogger.Logger.Error("Failed to parse the feed.", "Caused by", err)
		return nil, err
	}

	slogger.Logger.Info("Successfully collected the feed.", "Feed", feed.Title)
	return feed, nil
}
