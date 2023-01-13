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

// SetPrefix sets prefix in ip
func (ip *IPv6) SetPrefix(prefix string) {
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

// Random returns a random IPv6 address
func Random() *IPv6 {
	ip := &IPv6{}
	_, err := rand.Read(ip.b[:])
	if err != nil {
		log.Fatal(err)
	}

	return ip
}
