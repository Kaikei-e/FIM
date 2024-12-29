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

var dur = time.Duration(2) * time.Second
var safeurlClient = safeurl.GetConfigBuilder().
	EnableIPv6(true).
	SetTimeout(dur).
	SetAllowedPorts(80, 443).
	Build()

func TestCollectRSSFeed(t *testing.T) {
	var buf bytes.Buffer
	h := slog.NewJSONHandler(&buf, nil)
	slogger.Logger = slog.New(h)

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
				cl: safeurl.Client(safeurlClient),
			},
			want: want{
				feed: &gofeed.Feed{
					Title: "Lorem ipsum feed for an interval of 1 minutes with 10 item(s)",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid rss feed link",
			args: args{
				rssLink: url.URL{
					Scheme: "https",
					Host:   "invalid-url\\\\",
				},
				cl: safeurl.Client(safeurlClient),
			},
			want: want{
				feed: &gofeed.Feed{},
			},
			wantErr: true,
		},
		{
			name: "malicious rss feed link",
			args: args{
				rssLink: url.URL{
					Scheme: "https",
					Host:   "example.com",
					Path:   "/feed?param=<script>alert(1)</script>",
				},
				cl: safeurl.Client(safeurlClient),
			},
			want: want{
				feed: &gofeed.Feed{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CollectRSSFeed(tt.args.rssLink, tt.args.cl)
			if err != nil {
				t.Errorf("CollectRSSFeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CollectRSSFeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
