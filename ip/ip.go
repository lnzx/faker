package ip

import (
	"crypto/rand"
	"net"
)

func IPv4() string {
	ip := make([]byte, net.IPv4len)
	_, _ = rand.Read(ip)
	return net.IP(ip).String()
}
