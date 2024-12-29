package protector

import (
	"time"

	"github.com/doyensec/safeurl"
)

func InitClient() *safeurl.WrappedClient {

	dur := time.Duration(2) * time.Second
	config := safeurl.GetConfigBuilder().
		EnableIPv6(true).
		SetAllowedSchemes("https").
		SetTimeout(dur).
		SetAllowedPorts(80, 443).
		Build()

	safeurlClient := safeurl.Client(config)

	return safeurlClient
}
