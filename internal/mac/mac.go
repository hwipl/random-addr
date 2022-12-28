package mac

import (
	"crypto/rand"
	"fmt"
	"log"
)

// Random returns a random MAC address
func Random() string {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}

	// set local and unicast bits
	b[0] |= 0b00000010
	b[0] &= 0b11111110

	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
		b[0], b[1], b[2], b[3], b[4], b[5])
}
