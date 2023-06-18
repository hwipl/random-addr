package ipv4

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

// Type returns the type of ip
func (ip *IPv4) Type() string {
	if ip.Loopback() {
		return "loopback"
	}
	if ip.Unspecified() {
		return "unspecified"
	}
	pp := "public"
	if ip.Private() {
		pp = "private"
	}
	ubm := "unicast"
	if ip.Broadcast() {
		ubm = "broadcast"
	}
	if ip.Multicast() {
		ubm = "multicast"
	}
	return fmt.Sprintf("%s %s", pp, ubm)
}

// aaBracketTop returns the top part of an ascii art bracket with length l
func aaBracketTop(l int) string {
	if l < 0 {
		return ""
	}

	switch l {
	case 0:
		return ""
	case 1:
		return "|"
	case 2:
		return "|\\"
	case 3:
		return "|\\ "
	case 4:
		return " /\\ "
	}
	left := (l - 4) / 2
	right := l - 4 - left
	return " " + strings.Repeat("_", left) + "/\\" + strings.Repeat("_", right) + " "
}

// aaBracketBottom returns the bottom part of an ascii art bracket with length l
func aaBracketBottom(l int) string {
	if l < 0 {
		return ""
	}

	switch l {
	case 0:
		return ""
	case 1:
		return "|"
	case 2:
		return "||"
	}
	return "|" + strings.Repeat(" ", l-2) + "|"
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

	// get prefix bytes,
	// get number of bits to be overwritten
	b := p.Addr().As4()
	bits := p.Bits()

	// overwrite bits
	// try to overwrite full bytes first, then single bits
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

// Parse parses and returns the IPv4 address in s
func Parse(s string) *IPv4 {
	ip := &IPv4{}

	// parse ip with prefix
	if strings.Contains(s, "/") {
		p, err := netip.ParsePrefix(s)
		if err != nil {
			log.Fatal(err)
		}

		ip.b = p.Addr().As4()
		ip.pl = p.Bits()

		return ip
	}

	// parse ip without prefix
	a, err := netip.ParseAddr(s)
	if err != nil {
		log.Fatal(err)
	}

	ip.b = a.As4()

	return ip
}
