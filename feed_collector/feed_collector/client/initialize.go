package client

import "net/http"

func Initialize() http.Client {
	cl := http.Client{}

	return cl
}