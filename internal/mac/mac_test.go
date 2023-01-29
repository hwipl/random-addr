package mac

import (
	"testing"
)

// Random tests random MAC address creation
func TestRandom(t *testing.T) {
	mac := Random()
	if mac == nil {
		t.Errorf("invalid MAC")
	}
}

// TestRandomUI tests random universal individual address creation
func TestRandomUI(t *testing.T) {
	mac := RandomUI()
	if !mac.Universal() {
		t.Errorf("MAC not universal")
	}
	if !mac.Individual() {
		t.Errorf("MAC not individual")
	}
}

// TestRandomUU tests random universal unicast address creation
func TestRandomUU(t *testing.T) {
	mac := RandomUU()
	if !mac.Universal() {
		t.Errorf("MAC not universal")
	}
	if !mac.Unicast() {
		t.Errorf("MAC not unicast")
	}
}

// TestRandomUG tests random universal group address creation
func TestRandomUG(t *testing.T) {
	mac := RandomUG()
	if !mac.Universal() {
		t.Errorf("MAC not universal")
	}
	if !mac.Group() {
		t.Errorf("MAC not group")
	}
}

// TestRandomUM tests random universal multicast address creation
func TestRandomUM(t *testing.T) {
	mac := RandomUM()
	if !mac.Universal() {
		t.Errorf("MAC not universal")
	}
	if !mac.Multicast() {
		t.Errorf("MAC not multicast")
	}
}

// TestRandomLI tests random local individual address creation
func TestRandomLI(t *testing.T) {
	mac := RandomLI()
	if !mac.Local() {
		t.Errorf("MAC not local")
	}
	if !mac.Individual() {
		t.Errorf("MAC not individual")
	}
}

// TestRandomLU tests random local unicast address creation
func TestRandomLU(t *testing.T) {
	mac := RandomLU()
	if !mac.Local() {
		t.Errorf("MAC not local")
	}
	if !mac.Unicast() {
		t.Errorf("MAC not unicast")
	}
}

// TestRandomLG tests random local group address creation
func TestRandomLG(t *testing.T) {
	mac := RandomLG()
	if !mac.Local() {
		t.Errorf("MAC not local")
	}
	if !mac.Group() {
		t.Errorf("MAC not group")
	}
}

// TestRandomLM tests random local multicast address creation
func TestRandomLM(t *testing.T) {
	mac := RandomLM()
	if !mac.Local() {
		t.Errorf("MAC not local")
	}
	if !mac.Multicast() {
		t.Errorf("MAC not multicast")
	}
}

// TestParse tests Parse
func TestParse(t *testing.T) {
	// test with static mac
	want := "00:00:5e:00:53:01"
	got := Parse(want).String()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
	got = Parse("00-00-5e-00-53-01").String()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	// test with random mac
	want = Random().String()
	got = Parse(want).String()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
