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

// Decimal returns ip as dotted decimal
func (ip *IPv4) Decimal() string {
	return ip.Addr().String()
}

// String returns ip as String
func (ip *IPv4) String() string {
	return ip.Decimal()
}

// SetPrefix sets prefix in ip
func (ip *IPv4) SetPrefix(prefix string) {
	// parse prefix
	p, err := netip.ParsePrefix(prefix)
	if err != nil {
		log.Fatal(err)
	}

	// get prefix bytes
	b, err := p.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}

	// get number of bits to be overwritten
	bits := p.Bits()

	// overwrite bits
	// try to overwrite full bytes first, then single bits
	// ignore last byte of b because it's p.Bits()
	for i := 0; i < len(b)-1; i++ {
		if bits < 8 {
			// last byte, not full, overwrite bits
			ipBits := ip.b[i] & (0xff >> bits)
			bBits := b[i] & (0xff << (8 - bits))
			ip.b[i] = ipBits | bBits

			break
		}

		// full byte, overwrite byte
		ip.b[i] = b[i]

		bits -= 8
	}
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
