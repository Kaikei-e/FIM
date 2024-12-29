package protector

import (
	"net/url"
	"time"

	"github.com/doyensec/safeurl"
)

func AntiURLAttack(link string) (*url.URL, error) {
	ul, err := url.Parse(link)
	if err != nil {
		return nil, err
	}

	dur := time.Duration(2) * time.Second

	config := safeurl.GetConfigBuilder().
		EnableIPv6(true).
		SetTimeout(dur).
		SetAllowedPorts(80, 443).
		Build()

	cl := safeurl.Client(config)
	resp, err := cl.Get(link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ul, nil
}
