package rss_feed

import "time"

type Feed struct {
	Title         string
	Description   string
	URL           string
	Author        string
	PublishedDate time.Time
}
