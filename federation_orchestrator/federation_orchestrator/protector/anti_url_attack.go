package protector

import (
	"context"
	"errors"
	"net"
	"net/http"
	"net/url"
	"time"

	"federation_orchestrator/slogger"

	"github.com/doyensec/safeurl"
)

// AntiURLAttack prevents SSRF attack
func AntiURLAttack(link string, client *safeurl.WrappedClient) (*url.URL, error) {
	parsedURL, err := url.Parse(link)
	if err != nil {
		slogger.Logger.Error("failed to parse the URL", "URL", link, "Caused by", err)
		return nil, err
	}

	if parsedURL.Host == "" || parsedURL.Scheme == "" {
		slogger.Logger.Error("URL is not valid", "URL", link)
		return nil, errors.New("host or scheme is missing")
	}

	if parsedURL.RawQuery != "" || parsedURL.Fragment != "" {
		slogger.Logger.Error("URL contains query parameters or fragments", "URL", link)
		return nil, errors.New("URL contains query parameters or fragments")
	}

	ips, err := net.LookupIP(parsedURL.Hostname())
	if err != nil {
		slogger.Logger.Error("failed to solve host IP addr", "host", parsedURL.Hostname(), "caused by", err)
		return nil, err
	}

	for _, ip := range ips {
		if isPrivateIP(ip) {
			slogger.Logger.Error("URL contains private IP address", "URL", link, "IP", ip)
			return nil, errors.New("URL contains private IP address")
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		slogger.Logger.Error("Failed to make the GET request", "URL", link, "Caused by", err)
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		slogger.Logger.Error("Failed to get the response from the URL", "URL", link, "Caused by", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slogger.Logger.Error("URL returns non-200 status code", "URL", link, "status code", resp.StatusCode)
		return nil, errors.New("URL returns non-200 status code")
	}

	return parsedURL, nil
}

func isPrivateIP(ip net.IP) bool {
	privateIPBlocks := []net.IPNet{
		{
			IP:   net.IPv4(10, 0, 0, 0),
			Mask: net.CIDRMask(8, 32),
		},
		{
			IP:   net.IPv4(172, 16, 0, 0),
			Mask: net.CIDRMask(12, 32),
		},
		{
			IP:   net.IPv4(192, 168, 0, 0),
			Mask: net.CIDRMask(16, 32),
		},
		{
			IP:   net.ParseIP("::1"),
			Mask: net.CIDRMask(128, 128),
		},
		{
			IP:   net.ParseIP("fc00::"),
			Mask: net.CIDRMask(7, 128),
		},
	}

	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}
