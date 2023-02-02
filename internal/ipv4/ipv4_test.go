package ipv4

import "testing"

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
