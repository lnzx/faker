package ip

import (
	"crypto/rand"
	"net"
)

func IPv4() string {
	ip := make(net.IP, net.IPv4len)
	_, _ = rand.Read(ip)
	return ip.String()
}
