## RP
###### Reverse Proxy with Weighted Round Robin balancer.

> A reverse proxy server is an intermediate connection point positioned at a network’s edge. It receives initial HTTP connection requests, acting like the actual endpoint.

<img src="https://raw.githubusercontent.com/ahmdrz/rp/master/resources/reverse-proxy.jpg" width="100%"/>

Image credit: <a href="https://www.imperva.com/learn/performance/reverse-proxy/">What is Reverse Proxy</a>

### Installation

```
$ go get -u github.com/ahmdrz/rp
```

### CommandLine

Running reverse proxy without balancer:

```
$ rp --config rpconfig.yaml --verbose serve
```

Example of configuration file

```yaml
listenaddr: 0.0.0.0:8080

targets:
- address: http://api.server1.com
  weight: 3
- address: http://api.server2.com
  weight: 2
```

Generate default configuration file

```
$ rp --config rpconfig.yaml generate
```

### API

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
  proxy := rp.New()
  proxy.Log(true)
  
  // Add will append a new endpoint to rp
  // Round-Robin only works if you add more than 1 endpoint
  // weights must be positive and greater than 0
  proxy.Add(newURL("https://api.server1.com"), 1)
	
  proxy.ListenAndServe("0.0.0.0:8080")
}

```

### Using Docker

```dockerfile
FROM ahmdrz/rp:latest
COPY rpconfig.yaml .
EXPOSE 8080
CMD ["rp", "--verbose", "serve"]
```

### Todo

- [x] Better CLI Application
- [ ] Failover Solution
- [ ] Health Check
