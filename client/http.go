package client

import (
	"net/http"
	"sync"
	"time"
)

var httpClientOnce = sync.OnceValue(func() *http.Client {
	return &http.Client{
		Timeout: time.Duration(10) * time.Second,
	}
})

func GetDefaultHttpClient() *http.Client {
	return httpClientOnce()
}
