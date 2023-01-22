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

// Prefix returns ip as Prefix
func (ip *IPv4) Prefix() netip.Prefix {
	return netip.PrefixFrom(ip.Addr(), ip.pl)
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

// Network returns the network part of ip
func (ip *IPv4) Network() string {
	return ip.Prefix().Masked().Addr().String()
}

// Host returns the host part of ip
func (ip *IPv4) Host() string {
	// create temporary array with only host bits set
	b := [4]byte{}
	bits := ip.pl
	for i := 0; i < len(b); i++ {
		if bits >= bitsPerByte {
			// full byte, skip to next byte
			bits -= bitsPerByte
			continue
		}

		// skip remaining network bits, set host bits
		ipBits := ip.b[i] & (0xff >> bits)
		b[i] |= ipBits
		bits = 0
	}

	return netip.AddrFrom4(b).String()
}

// Loopback returns wether ip is a loopback address
func (ip *IPv4) Loopback() bool {
	return ip.Addr().IsLoopback()
}

// Private returns wether ip is a private address
func (ip *IPv4) Private() bool {
	return ip.Addr().IsPrivate()
}

// Unspecified returns wether ip is the unspecified address
func (ip *IPv4) Unspecified() bool {
	return ip.Addr().IsUnspecified()
}

// Multicast returns wether ip is a multicast address
func (ip *IPv4) Multicast() bool {
	return ip.Addr().IsMulticast()
}

// Broadcast returns wether ip is a broadcast address
func (ip *IPv4) Broadcast() bool {
	// skip prefix bits
	bits := ip.pl
	for i := 0; i < len(ip.b); i++ {
		if bits >= bitsPerByte {
			// full byte, skip to next byte
			bits -= bitsPerByte
			continue
		}

		// skip remaining network bits, check if host bits are set to 1
		ipBits := ip.b[i] & (0xff >> bits)
		if ipBits != (0xff >> bits) {
			return false
		}
		bits = 0
	}
	return true
}

// Unicast returns wether ip is a unicast address
func (ip *IPv4) Unicast() bool {
	return !ip.Multicast() && !ip.Broadcast()
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

	// set new prefix length
	ip.pl = p.Bits()
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
