package ipv6

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/netip"
	"strings"
)

const (
	// bitsPerByte is the number of bits per byte
	bitsPerByte = 8
)

// IPv6 is an IPv6 address
type IPv6 struct {
	b [16]byte

	// pl is the prefix length
	pl int
}

// Addr returns ip as Addr
func (ip *IPv6) Addr() netip.Addr {
	return netip.AddrFrom16(ip.b)
}

// Prefix returns ip as Prefix
func (ip *IPv6) Prefix() netip.Prefix {
	return netip.PrefixFrom(ip.Addr(), ip.pl)
}

// Binary returns ip as a binary string
func (ip *IPv6) Binary() string {
	return fmt.Sprintf("%08b%08b:%08b%08b:"+
		"%08b%08b:%08b%08b:"+
		"%08b%08b:%08b%08b:"+
		"%08b%08b:%08b%08b",
		ip.b[0], ip.b[1], ip.b[2], ip.b[3],
		ip.b[4], ip.b[5], ip.b[6], ip.b[7],
		ip.b[8], ip.b[9], ip.b[10], ip.b[11],
		ip.b[12], ip.b[13], ip.b[14], ip.b[15],
	)
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

	// get prefix bytes,
	// get number of bits to be overwritten
	b := p.Addr().As16()
	bits := p.Bits()

	// overwrite bits
	// try to overwrite full bytes first, then single bits
	// ignore last byte of b because it's p.Bits()
	for i := 0; i < len(b); i++ {
		if bits < bitsPerByte {
			// last byte, not full, overwrite bits
			ipBits := ip.b[i] & (0xff >> bits)
			bBits := b[i] & (0xff << (bitsPerByte - bits))
			ip.b[i] = ipBits | bBits

			break
		}

		// full byte, overwrite byte
		ip.b[i] = b[i]

		bits -= bitsPerByte
	}

	// set new prefix length
	ip.pl = p.Bits()
}

// SetPrefixLength sets prefix length of ip in number of bits
func (ip *IPv6) SetPrefixLength(numBits int) {
	ip.pl = numBits
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

// Parse parses and returns the IPv6 address in s
func Parse(s string) *IPv6 {
	ip := &IPv6{}

	// parse ip with prefix
	if strings.Contains(s, "/") {
		p, err := netip.ParsePrefix(s)
		if err != nil {
			log.Fatal(err)
		}

		ip.b = p.Addr().As16()
		ip.pl = p.Bits()

		return ip
	}

	// parse ip without prefix
	a, err := netip.ParseAddr(s)
	if err != nil {
		log.Fatal(err)
	}

	ip.b = a.As16()

	return ip
}
