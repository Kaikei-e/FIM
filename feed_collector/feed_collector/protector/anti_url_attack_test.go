package protector

import (
	logger "feed_collector/slogger"
	"log/slog"
	"net/url"
	"os"
	"reflect"
	"testing"
)

func TestAntiURLAttack(t *testing.T) {
	// Mock the logger
	mockLogger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	originalLogger := logger.Logger
	logger.Logger = mockLogger
	defer func() { logger.Logger = originalLogger }()

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
			got, err := AntiURLAttack(tt.args.link)
			if err != nil {
				t.Errorf("AntiURLAttack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want.url) {
				t.Errorf("AntiURLAttack() = %v, want %v", got, tt.want.url)
			}
			if (got != url.URL{}) && got.Host != tt.want.host {
				t.Errorf("AntiURLAttack() = %v, want %v", got.Host, tt.want.host)
			}
		})
	}
}
