package rp

import (
	"sync"
)

type roundRobin struct {
	sync.Mutex
	conns         []*proxyConnection
	idx           int
	maxWeight     int
	currentGCD    int
	currentWeight int
}

func newRoundRobin() *roundRobin {
	out := &roundRobin{
		conns: make([]*proxyConnection, 0),
		idx:   -1,
	}
	return out
}

func (r *roundRobin) Add(p *proxyConnection) {
	if p.weight <= 0 {
		panic("weight must be greater than 0")
	}
	if r.currentGCD == 0 {
		r.currentGCD = p.weight
		r.maxWeight = p.weight
	} else {
		r.currentGCD = gcd(r.currentGCD, p.weight)
		if r.maxWeight < p.weight {
			r.maxWeight = p.weight
		}
	}
	r.conns = append(r.conns, p)
}

func (r *roundRobin) Get() *proxyConnection {
	r.Lock()
	defer r.Unlock()

	if len(r.conns) == 1 {
		return r.conns[0]
	}

	for {
		r.idx = (r.idx + 1) % len(r.conns)
		if r.idx == 0 {
			r.currentWeight = r.currentWeight - r.currentGCD
			if r.currentWeight <= 0 {
				r.currentWeight = r.maxWeight
			}
		}

		if r.conns[r.idx].weight >= r.currentWeight {
			return r.conns[r.idx]
		}
	}
}

func gcd(x, y int) int {
	var t int
	for {
		t = (x % y)
		if t > 0 {
			x = y
			y = t
		} else {
			return y
		}
	}
}
