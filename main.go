package main

import (
	"flag"
	"log"
	"net/url"

	rp "github.com/ahmdrz/rp/reverse-proxy"
)

func newURL(addr string) *url.URL {
	u, err := url.Parse(addr)
	if err != nil {
		log.Fatal(err)
	}
	return u
}

func main() {
	var (
		remote  string
		bind    string
		logMode bool
	)
	flag.StringVar(&bind, "l", "0.0.0.0:8888", "listen on ip:port")
	flag.StringVar(&remote, "r", "http://google.com", "reverse proxy target")
	flag.BoolVar(&logMode, "v", false, "log mode")
	flag.Parse()

	proxy := rp.New()
	proxy.Log(logMode)
	proxy.Add(newURL(remote), 1)
	proxy.ListenAndServe(bind)
}
