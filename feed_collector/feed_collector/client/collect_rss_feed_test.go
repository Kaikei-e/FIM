package client

import (
	"bytes"
	"feed_collector/slogger"
	"log/slog"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/doyensec/safeurl"
	"github.com/mmcdole/gofeed"
)

func TestCollectRSSFeed(t *testing.T) {
	// Mock the logger
	var buf bytes.Buffer
	h := slog.NewJSONHandler(&buf, nil)
	slogger.Logger = slog.New(h)
	defer func() { slogger.Logger = nil }()

	// Mock the safeurl client
	dur := time.Duration(2) * time.Second
	config := safeurl.GetConfigBuilder().
		EnableIPv6(true).
		SetAllowedSchemes("https").
		SetTimeout(dur).
		SetAllowedPorts(80, 443).
		Build()
	safeurlClient := safeurl.Client(config)

	type args struct {
		rssLink url.URL
		cl      *safeurl.WrappedClient
	}
	type want struct {
		feed *gofeed.Feed
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "collect rss feed",
			args: args{
				rssLink: url.URL{
					Scheme: "https",
					Host:   "lorem-rss.herokuapp.com",
					Path:   "/feed",
				},
				cl: safeurlClient,
			},
			want: want{
				feed: &gofeed.Feed{
					Title: "Lorem ipsum feed for an interval of 1 minutes with 10 item(s)",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CollectRSSFeed(tt.args.rssLink, tt.args.cl)
			if (err != nil) != tt.wantErr {
				t.Errorf("CollectRSSFeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Title, tt.want.feed.Title) {
				t.Errorf("CollectRSSFeed() = %v, want %v", got.Title, tt.want.feed.Title)
			}
		})
	}
}
