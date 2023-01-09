package ipv4

import (
	"crypto/rand"
	"log"
	"net/netip"
)

// IPv4 is an IPv4 address
type IPv4 struct {
	b [4]byte
}

// Addr returns ip as Addr
func (ip *IPv4) Addr() netip.Addr {
	return netip.AddrFrom4(ip.b)
}

// String returns ip as String
func (ip *IPv4) String() string {
	return ip.Addr().String()
}

// Random returns a random IPv4 address
func Random() *IPv4 {
	ip := &IPv4{}
	_, err := rand.Read(ip.b[:])
	if err != nil {
		log.Fatal(err)
	}

	return ip
}
