package protector

import (
	logger "feed_collector/slogger"
	"log/slog"
	"net/url"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/doyensec/safeurl"
)

func TestAntiURLAttack(t *testing.T) {
	// Mock the logger
	mockLogger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	originalLogger := logger.Logger
	logger.Logger = mockLogger
	defer func() { logger.Logger = originalLogger }()

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
		link string
	}
	type want struct {
		url  *url.URL
		host string
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "good url",
			args: args{
				link: "https://example.com",
			},
			want: want{
				url:  &url.URL{Scheme: "https", Host: "example.com"},
				host: "example.com",
			},
			wantErr: false,
		},
		{
			name: "bad url",
			args: args{
				link: "https://example.com/feed?param=<script>alert(1)</script>",
			},
			want: want{
				url:  nil,
				host: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AntiURLAttack(tt.args.link, safeurlClient)
			if (err != nil) != tt.wantErr {
				t.Errorf("AntiURLAttack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want.url) {
				t.Errorf("AntiURLAttack() = %v, want %v", got, tt.want.url)
			}
			if got != nil && got.Host != tt.want.host {
				t.Errorf("AntiURLAttack() = %v, want %v", got.Host, tt.want.host)
			}
		})
	}
}
