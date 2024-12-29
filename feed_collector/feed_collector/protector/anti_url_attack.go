package protector

import (
	"errors"
	"feed_collector/slogger"
	"net/url"

	"github.com/doyensec/safeurl"
)

func AntiURLAttack(link string, client *safeurl.WrappedClient) (*url.URL, error) {
	parsedURL, err := url.Parse(link)
	if err != nil {
		slogger.Logger.Error("Failed to parse the URL.", "Caused by", err)
		return nil, err
	}

	if parsedURL.Host == "" || parsedURL.Scheme == "" {
		slogger.Logger.Error("Invalid URL.", "URL", link)
		return nil, errors.New("host is empty and url is not valid")
	}

	if parsedURL.RawQuery != "" {
		slogger.Logger.Error("URL contains query parameters.", "URL", link)
		return nil, errors.New("url contains query parameters")
	}

	resp, err := client.Get(parsedURL.String())
	if err != nil {
		slogger.Logger.Error("Failed to get the URL.", "Caused by", err)
		return nil, err
	}
	defer resp.Body.Close()

	return parsedURL, nil
}
