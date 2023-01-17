package ipv4

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/netip"
)

const (
	// bitsPerByte is the number of bits per byte
	bitsPerByte = 8
)

// IPv4 is an IPv4 address
type IPv4 struct {
	b [4]byte

	// pl is the prefix length
	pl int
}

// Addr returns ip as Addr
func (ip *IPv4) Addr() netip.Addr {
	return netip.AddrFrom4(ip.b)
}

// Decimal returns ip as dotted decimal
func (ip *IPv4) Decimal() string {
	return ip.Addr().String()
}

// Binary returns ip as a binary string
func (ip *IPv4) Binary() string {
	return fmt.Sprintf("%08b.%08b.%08b.%08b",
		ip.b[0], ip.b[1], ip.b[2], ip.b[3])
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
}

// SetPrefixLength sets prefix length of ip in number of bits
func (ip *IPv4) SetPrefixLength(numBits int) {
	ip.pl = numBits
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
