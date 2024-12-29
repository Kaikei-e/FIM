package client

import (
	"feed_collector/protector"
	"log"
	"net/url"

	"github.com/doyensec/safeurl"
	"github.com/mmcdole/gofeed"
)

func CollectRSSFeed(rssLink url.URL, cl *safeurl.WrappedClient) (*gofeed.Feed, error) {
	targetLink, err := protector.AntiURLAttack(rssLink.String())
	if err != nil {
		log.Println("This link is potentially malicious")
		return nil, err
	}

	fp := gofeed.NewParser()

	feed, err := fp.ParseURL(targetLink.String())
	if err != nil {
		return nil, err
	}

	return feed, nil
}
