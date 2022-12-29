package mac

import (
	"crypto/rand"
	"fmt"
	"log"
)

// MAC is a MAC address
type MAC struct {
	b [6]byte
}

// Hex returns MAC as hex string
func (m *MAC) Hex() string {
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
		m.b[0], m.b[1], m.b[2], m.b[3], m.b[4], m.b[5])
}

// Binary returns MAC as a binary string
func (m *MAC) Binary() string {
	return fmt.Sprintf("%08b:%08b:%08b:%08b:%08b:%08b",
		m.b[0], m.b[1], m.b[2], m.b[3], m.b[4], m.b[5])
}

// String returns MAC as a string
func (m *MAC) String() string {
	return m.Hex()
}

// Random returns a random MAC address
func Random() *MAC {
	m := &MAC{}
	_, err := rand.Read(m.b[:])
	if err != nil {
		log.Fatal(err)
	}

	// set local and unicast bits
	m.b[0] |= 0b00000010
	m.b[0] &= 0b11111110

	return m
}
