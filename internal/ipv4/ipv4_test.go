package ipv4

import "testing"

// TestType tests Type of IPv4
func TestType(t *testing.T) {
	// test loopback
	loop := Parse("127.0.0.1")
	want := "loopback"
	got := loop.Type()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	// test unspecified
	unspec := Parse("0.0.0.0")
	want = "unspecified"
	got = unspec.Type()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	// test public unicast
	pubUC := Parse("1.2.3.4")
	want = "public unicast"
	got = pubUC.Type()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	// test public broadcast
	pubBC := Parse("1.2.3.255/24")
	want = "public broadcast"
	got = pubBC.Type()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	// test public multicast
	pubMC := Parse("224.0.0.1/24")
	want = "public multicast"
	got = pubMC.Type()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	// test private unicast
	privUC := Parse("10.2.3.4")
	want = "private unicast"
	got = privUC.Type()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	// test private broadcast
	privBC := Parse("10.2.3.255/24")
	want = "private broadcast"
	got = privBC.Type()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

// TestSetPrefix tests SetPrefix of IPv4
func TestSetPrefix(t *testing.T) {
	ip := Random()
	ip.SetPrefix("127.0.0.0/8")
	want := "127.0.0.0"
	got := ip.Network()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

// TestSetPrefixLength tests SetPrefixLength of IPv4
func TestSetPrefixLength(t *testing.T) {
	ip := Parse("127.0.0.1")
	ip.SetPrefixLength(8)
	want := "127.0.0.0"
	got := ip.Network()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

// TestRandom tests creation of random IPv4 addresses
func TestRandom(t *testing.T) {
	ip := Random()
	if ip == nil {
		t.Errorf("invalid IPv4 address")
	}
}

// TestParse tests parsing of IPv4 addresses
func TestParse(t *testing.T) {
	// test without prefix
	want := "127.0.0.1"
	got := Parse(want).Addr().String()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	// test with prefix
	want = "127.0.0.1/8"
	got = Parse(want).Prefix().String()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
