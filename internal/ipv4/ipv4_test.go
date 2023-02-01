package ipv4

import "testing"

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
