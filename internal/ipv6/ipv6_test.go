package ipv6

import "testing"

// TestRandom tests random IPv6 creation
func TestRandom(t *testing.T) {
	ip := Random()
	if ip == nil {
		t.Errorf("invalid IPv6 address")
	}
}

// TestParse tests IPv6 address parsing
func TestParse(t *testing.T) {
	// test without prefix
	want := "fe80::1"
	got := Parse(want).Addr().String()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	// test with prefix
	want = "fe80::1/64"
	got = Parse(want).Prefix().String()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
