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
$ rp -l "localhost:8080" -r "https://api.server1.com" -v true
```

Running reverse proxy with balancer:

```
$ rp -l "localhost:8080" -r "https://api.server1.com,https://api.server2.com" -v true
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

```
$ docker pull ahmdrz/rp:latest
$ docker run -p 8888:8888 -e REMOTE="https://api.server1.com,https://api.server2.com" ahmdrz/rp:latest
```


### Todo

- [ ] Failover Solution
- [ ] Health Check
