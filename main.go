package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"strings"

	rp "github.com/ahmdrz/rp/reverse-proxy"
)

func main() {
	var (
		remote  string
		bind    string
		logMode bool
	)
	flag.StringVar(&bind, "l", "0.0.0.0:8888", "listen on ip:port")
	flag.StringVar(&remote, "r", "", "reverse proxy target")
	flag.BoolVar(&logMode, "v", false, "log mode")
	flag.Parse()

	if remote == "" {
		remote = os.Getenv("REMOTE")
	}

	if !logMode {
		logMode = os.Getenv("VERBOSE") == "true"
	}

	if os.Getenv("LISTEN_ADDR") != "" {
		bind = os.Getenv("LISTEN_ADDR")
	}

	proxy := rp.New()
	proxy.Log(logMode)
	for _, target := range strings.Split(remote, ",") {
		if target == "" {
			log.Fatal("[bad target] one of targets is empty")
		}
		t, err := url.Parse(target)
		if err != nil {
			log.Fatal("[bad target]", err)
		}

		// TODO: get weight from cli/env
		proxy.Add(t, 1)
	}

	log.Printf("Starting reverse-proxy on %s ...", bind)
	log.Fatal(proxy.ListenAndServe(bind))
}
