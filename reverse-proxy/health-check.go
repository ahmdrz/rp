package rp

import (
	"net"
	"net/http"
	"net/url"
	"time"
)

type HealthCheck struct {
	endpoint  *url.URL
	attemps   uint
	client    *http.Client
	sleepTime time.Duration
}

func NewHealthCheck(endpoint *url.URL, attemps uint) *HealthCheck {
	return &HealthCheck{
		endpoint:  endpoint,
		attemps:   attemps,
		sleepTime: 1 * time.Second,
		client: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				Dial: (&net.Dialer{
					Timeout:   10 * time.Second,
					KeepAlive: 0,
				}).Dial,
				TLSHandshakeTimeout:   5 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				MaxIdleConnsPerHost:   1,
				DisableCompression:    true,
				DisableKeepAlives:     true,
				ResponseHeaderTimeout: 5 * time.Second,
			},
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Timeout: 10 * time.Second,
		},
	}
}

func (h *HealthCheck) check() bool {
	req, err := http.NewRequest("GET", h.endpoint.String(), nil)
	if err != nil {
		return false
	}
	for i := uint(0); i < h.attemps; i++ {
		resp, err := h.client.Do(req)
		if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return true
		}
		time.Sleep(h.sleepTime)
	}
	return false
}
