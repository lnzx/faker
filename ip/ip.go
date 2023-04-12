package ip

import (
	randc "crypto/rand"
	"math/rand"
	"net"
)

type ipOptions struct {
	bogon bool
}

type Option interface {
	apply(o *ipOptions)
}

type funcOption func(*ipOptions)

func (f funcOption) apply(o *ipOptions) {
	f(o)
}

func WithBogon() Option {
	return funcOption(func(o *ipOptions) {
		o.bogon = true
	})
}

func IPv4(options ...Option) string {
	opts := ipOptions{}
	for _, opt := range options {
		opt.apply(&opts)
	}
	ip := make(net.IP, net.IPv4len)
	_, _ = randc.Read(ip)
	if !opts.bogon {
		ip[0] = byte(1 + rand.Intn(223))
	}
	return ip.String()
}
