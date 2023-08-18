package rp

import (
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type proxyConnection struct {
	reverseProxy *httputil.ReverseProxy
	target       *url.URL
	weight       int
	healthCheck  *HealthCheck
	isHealthy    bool
}

func newProxyConnection(target *url.URL, weight int, healthCheck *HealthCheck) *proxyConnection {
	return &proxyConnection{
		reverseProxy: httputil.NewSingleHostReverseProxy(target),
		target:       target,
		weight:       weight,
		isHealthy:    true,
		healthCheck:  healthCheck,
	}
}

func (p *proxyConnection) String() string {
	return p.target.String()
}

func (p *proxyConnection) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.URL.Host = p.target.Host
	r.URL.Scheme = p.target.Scheme
	r.Host = p.target.Host
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	p.reverseProxy.Transport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 10 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          10,
		IdleConnTimeout:       30 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	p.reverseProxy.ServeHTTP(w, r)
}
