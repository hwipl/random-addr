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

// Hex returns ip as a hexadecimal string
func (ip *IPv6) Hex() string {
	return ip.Addr().String()
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

// Network returns the network part of ip as string
func (ip *IPv6) Network() string {
	return ip.Prefix().Masked().Addr().String()
}

// Subnet returns the subnet part of ip as string
func (ip *IPv6) Subnet() string {
	bits := ip.pl
	if bits >= 64 {
		// no subnet bits
		return ""
	}

	// create temporary array with only subnet bits set
	b := [16]byte{}
	for i := 0; i < 8; i++ {
		if bits >= bitsPerByte {
			// full byte, skip to next byte
			bits -= bitsPerByte
			continue
		}

		// skip remaining prefix bits, set subnet bits
		ipBits := ip.b[i] & (0xff >> bits)
		b[i] |= ipBits
		bits = 0
	}

	return netip.AddrFrom16(b).String()
}

// IID returns the interface identifier of ip as string
func (ip *IPv6) IID() string {
	return fmt.Sprintf("%02x%02x:%02x%02x:"+
		"%02x%02x:%02x%02x",
		ip.b[8], ip.b[9], ip.b[10], ip.b[11],
		ip.b[12], ip.b[13], ip.b[14], ip.b[15],
	)
}

// IPv4Mapped returns wether ip is an IPv4-mapped IPv6 address
func (ip *IPv6) IPv4Mapped() bool {
	return ip.Addr().Is4In6()
}

// GlobalUnicast returns wether ip is a global unicast address
func (ip *IPv6) GlobalUnicast() bool {
	return ip.Addr().IsGlobalUnicast()
}

// InterfaceLocalMulticast returns wether ip is an interface local
// multicast address
func (ip *IPv6) InterfaceLocalMulticast() bool {
	return ip.Addr().IsInterfaceLocalMulticast()
}

// LinkLocalMulticast returns wether ip is a local multicast address
func (ip *IPv6) LinkLocalMulticast() bool {
	return ip.Addr().IsLinkLocalMulticast()
}

// LinkLocalUnicast returns wether ip is a local unicast address
func (ip *IPv6) LinkLocalUnicast() bool {
	return ip.Addr().IsLinkLocalUnicast()
}

// Loopback returns wether ip is a loopback address
func (ip *IPv6) Loopback() bool {
	return ip.Addr().IsLoopback()
}

// Multicast returns wether ip is a multicast address
func (ip *IPv6) Multicast() bool {
	return ip.Addr().IsMulticast()
}

// Unicast returns wether ip is a unicast address
func (ip *IPv6) Unicast() bool {
	return !ip.Multicast()
}

// Private returns wether ip is a private address
func (ip *IPv6) Private() bool {
	return ip.Addr().IsPrivate()
}

// Unspecified returns wether ip is the unspecified address
func (ip *IPv6) Unspecified() bool {
	return ip.Addr().IsUnspecified()
}

// Type return the type of ip
func (ip *IPv6) Type() string {
	if ip.Loopback() {
		return "loopback"
	}
	if ip.Unspecified() {
		return "unspecified"
	}
	if ip.GlobalUnicast() {
		return "global unicast"
	}
	if ip.LinkLocalMulticast() {
		return "link-local multicast"
	}
	if ip.LinkLocalUnicast() {
		return "link-local unicast"
	}
	if ip.IPv4Mapped() {
		return "IPv4-mapped"
	}
	pp := "public"
	if ip.Private() {
		pp = "private"
	}
	um := "unicast"
	if ip.Multicast() {
		um = "multicast"
	}
	return fmt.Sprintf("%s %s", pp, um)
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

// ExplainBin returns an explanation of the IP and its structure as string
func (ip *IPv6) ExplainBin() string {
	// consider up to 3 dots in 32 bit address,
	// calculate prefix and host length for ascii art bracket creation
	pl := ip.pl + (ip.pl / (2 * bitsPerByte))
	sl := 67 - pl
	il := 67
	skip := ""
	if ip.pl != 0 && ip.pl%(2*bitsPerByte) == 0 {
		// prefix ends exacly at a dot,
		// omit this dot in ascii art bracket for prefix,
		// set skip to start ascii art bracket for host after this dot
		pl -= 1
		skip = " "
	}

	return fmt.Sprintf(`Network: %s    Subnet: %s    IID: %s
      %s%s%s %s
      %s%s%s %s
Bin:  %s
Type: %s
`,
		ip.Network(), ip.Subnet(), ip.IID(),
		aaBracketTop(pl), skip, aaBracketTop(sl), aaBracketTop(il),
		aaBracketBottom(pl), skip, aaBracketBottom(sl), aaBracketBottom(il),
		ip.Binary(),
		ip.Type(),
	)
}

// String returns ip as String
func (ip *IPv6) String() string {
	return ip.Hex()
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
