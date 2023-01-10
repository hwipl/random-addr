package ipv6

import (
	"crypto/rand"
	"log"
	"net/netip"
)

// IPv6 is an IPv6 address
type IPv6 struct {
	b [16]byte
}

// Addr returns ip as Addr
func (ip *IPv6) Addr() netip.Addr {
	return netip.AddrFrom16(ip.b)
}

// String returns ip as String
func (ip *IPv6) String() string {
	return ip.Addr().String()
}

// Random returns a random IPv6 address
func Random() *IPv6 {
	ip := &IPv6{}
	_, err := rand.Read(ip.b[:])
	if err != nil {
		log.Fatal(err)
	}

	return ip
}
