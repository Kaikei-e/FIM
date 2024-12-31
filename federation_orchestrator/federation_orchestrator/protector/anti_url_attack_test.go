package protector

import (
	"federation_orchestrator/slogger"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"
	"time"

	"log/slog"

	"github.com/doyensec/safeurl"
)

func MockLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, nil))
}

func TestAntiURLAttack(t *testing.T) {
	mockLogger := MockLogger()
	originalLogger := slogger.Logger
	slogger.Logger = mockLogger
	defer func() { slogger.Logger = originalLogger }()

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/valid-feed":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Valid Feed Content"))
		case "/not-found-feed":
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not Found"))
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
		}
	}))
	defer testServer.Close()

	dur := 5 * time.Second
	clConfig := safeurl.GetConfigBuilder().
		EnableIPv6(true).
		SetAllowedSchemes("https").
		SetTimeout(dur).
		SetAllowedPorts(80, 443).
		Build()

	cl := safeurl.Client(clConfig)

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
			name: "有効なURL",
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
			name: "Invalid URL - Banned to have query params",
			args: args{
				link: "https://example.com/feed?param=<script>alert(1)</script>",
			},
			want: want{
				url:  nil,
				host: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid URL - No internal IP addr",
			args: args{
				link: "https://127.0.0.1/attack",
			},
			want: want{
				url:  nil,
				host: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid URL - Not allowed scheme",
			args: args{
				link: "ftp://example.com/feed",
			},
			want: want{
				url:  nil,
				host: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid URL - No query parameters but 404",
			args: args{
				link: "https://example.com/not-found-feed",
			},
			want: want{
				url:  nil,
				host: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid URL - Contains fragment",
			args: args{
				link: "https://example.com/valid-feed#section",
			},
			want: want{
				url:  nil,
				host: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid URL - No host",
			args: args{
				link: "https:///invalid-host",
			},
			want: want{
				url:  nil,
				host: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid URL - Empty string",
			args: args{
				link: "",
			},
			want: want{
				url:  nil,
				host: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid URL - Parsing error",
			args: args{
				link: "://invalid-url",
			},
			want: want{
				url:  nil,
				host: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := AntiURLAttack(tt.args.link, cl)
			if (err != nil) != tt.wantErr {
				t.Errorf("AntiURLAttack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want.url) {
				t.Errorf("AntiURLAttack() got = %v, want = %v", got, tt.want.url)
			}
			if got != nil && got.Host != tt.want.host {
				t.Errorf("AntiURLAttack() got.Host = %v, want = %v", got.Host, tt.want.host)
			}
		})
	}
}
