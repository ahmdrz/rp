# RP

> Reverse-Proxy (rp) with Weighted Round Robin load-balancer.

---

### Installation

```
$ go get -u github.com/ahmdrz/rp
```

### CommandLine

```
$ rp -l "localhost:8080" -r "https://time.ir" -v true
```

### Example

```go
package main

import (
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
  target := "https://time.ir"
  weight := 1
  
  proxy := rp.New()
	proxy.Log(true)
  
	// Add will append a new endpoint to rp
  // Round-Robin only works if you add more than 1 endpoint
  // weights must be positive and greater than 0
  proxy.Add(newURL(target), weight)
	
  proxy.ListenAndServe("0.0.0.0:8080")
}

```
