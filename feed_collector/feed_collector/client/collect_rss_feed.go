package client

import (
	"feed_collector/protector"
	"feed_collector/slogger"
	"net/url"

	"github.com/doyensec/safeurl"
	"github.com/mmcdole/gofeed"
)

func CollectRSSFeed(rssLink url.URL, cl *safeurl.WrappedClient) (*gofeed.Feed, error) {
	linkPath := rssLink.Path
	rssLink.Path = ""

	targetLink, err := protector.AntiURLAttack(rssLink.String(), cl)
	if err != nil {
		slogger.Logger.Error("Failed to protect against URL attack.", "Caused by", err)
		return nil, err
	}

	fromattedPath, err := url.JoinPath(targetLink.String(), linkPath)
	if err != nil {
		slogger.Logger.Error("Failed to join the path.", "Caused by", err)
		return nil, err
	}

	reParsedURL, err := url.Parse(fromattedPath)
	if err != nil {
		slogger.Logger.Error("Failed to parse the URL.", "Caused by", err)
		return nil, err
	}

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(reParsedURL.String())
	if err != nil {
		slogger.Logger.Error("Failed to parse the feed.", "Caused by", err)
		return nil, err
	}

	slogger.Logger.Info("Successfully collected the feed.", "Feed", feed.Title)
	return feed, nil
}
