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

// Universal returns true if MAC is globally unique,
// i.e. Universal/Local (U/L) bit is 0
func (m *MAC) Universal() bool {
	return m.b[0]&0b00000010 == 0
}

// Local returns true if MAC is locally administered,
// i.e. Universal/Local (U/L) bit is 1
func (m *MAC) Local() bool {
	return !m.Universal()
}

// UL returns the U/L bit as string
func (m *MAC) UL() string {
	if m.Local() {
		return "Local"
	}
	return "Universal"
}

// Individual returns true if MAC's Individual/Group (I/G) bit is 0
func (m *MAC) Individual() bool {
	return m.b[0]&0b00000001 == 0
}

// Group returns true if MAC's Individual/Group (I/G) bit is 1
func (m *MAC) Group() bool {
	return !m.Individual()
}

// Unicast returns true if MAC is unicast,
// i.e. Individual/Group (I/G) bit is 0
func (m *MAC) Unicast() bool {
	return m.Individual()
}

// Multicast returns true if MAC is multicast,
// i.e. Individual/Group (I/G) bit is 1
func (m *MAC) Multicast() bool {
	return m.Group()
}

// IG returns the I/G (Unicast/Multicast) bit as a string
func (m *MAC) IG() string {
	if m.Group() {
		return "Group (Multicast)"
	}
	return "Individual (Unicast)"
}

// OUI returns the OUI part of MAC as string
func (m *MAC) OUI() string {
	return fmt.Sprintf("%02x:%02x:%02x", m.b[0], m.b[1], m.b[2])
}

// NIC returns the NIC specific part of MAC as string
func (m *MAC) NIC() string {
	return fmt.Sprintf("%02x:%02x:%02x", m.b[3], m.b[4], m.b[5])
}

// Explain returns an explanation of the MAC and its structure as string
func (m *MAC) Explain() string {
	return fmt.Sprintf(`           OUI: %s          NIC specific: %s
      ___________/\___________   ___________/\___________
     |                        | |                        |
Hex:    %02x   :   %02x   :   %02x   :   %02x   :   %02x   :   %02x
Bin: %s
           ||
           ||_ I/G: %s
           |__ U/L: %s`,
		m.OUI(), m.NIC(),
		m.b[0], m.b[1], m.b[2], m.b[3], m.b[4], m.b[5],
		m.Binary(),
		m.IG(), m.UL(),
	)
}

// ExplainHex returns an explanation of the hex representation of the MAC
// and its structure as string
func (m *MAC) ExplainHex() string {
	return fmt.Sprintf(`    OUI     NIC specific
 ___/\___   ___/\___
|        | |        |
 %s : %s
        |
        |_ I/G: %s
        |_ U/L: %s`,
		m.OUI(), m.NIC(), m.IG(), m.UL(),
	)
}

// ExplainBin returns an explanation of the binary representation of the MAC
// and its structure as string
func (m *MAC) ExplainBin() string {
	return fmt.Sprintf(`      OUI: %s          NIC specific: %s
 ___________/\___________   ___________/\___________
|                        | |                        |
%s
      ||
      ||_ I/G: %s
      |__ U/L: %s`,
		m.OUI(), m.NIC(),
		m.Binary(),
		m.IG(), m.UL(),
	)
}

// All returns all information about the MAC as string
func (m *MAC) All() string {
	return fmt.Sprintf(`Hex:          %s
OUI:          %s
NIC specific: %s
Binary:       %s
U/L:          %s
I/G:          %s`,
		m.Hex(),
		m.OUI(),
		m.NIC(),
		m.Binary(),
		m.UL(),
		m.IG(),
	)
}

// Table returns all information about the MAC as a table in a string
func (m *MAC) Table() string {
	return fmt.Sprintf(
		` ----------------------------------------------------------------------
| Hex          | %-53s |
| OUI          | %-53s |
| NIC specific | %-53s |
| Binary       | %-53s |
| U/L          | %-53s |
| I/G          | %-53s |
 ----------------------------------------------------------------------`,
		m.Hex(),
		m.OUI(),
		m.NIC(),
		m.Binary(),
		m.UL(),
		m.IG(),
	)
}

// SetUniversal sets the U/L bit of the MAC address to 0 (universal)
func (m *MAC) SetUniversal() {
	m.b[0] &= 0b11111101
}

// SetLocal sets the U/L bit of the MAC address to 1 (local)
func (m *MAC) SetLocal() {
	m.b[0] |= 0b00000010
}

// SetUL sets if the address is universal via the U/L bit
func (m *MAC) SetUL(universal bool) {
	if universal {
		m.SetUniversal()
		return
	}
	m.SetLocal()
}

// SetIndividual sets the I/G bit of the MAC address to 0 (individual)
func (m *MAC) SetIndividual() {
	m.b[0] &= 0b11111110
}

// SetGroup sets the I/G bit of the MAC address to 1 (group)
func (m *MAC) SetGroup() {
	m.b[0] |= 0b00000001
}

// SetUnicast sets I/G bit of the MAC address to unicast (individual)
func (m *MAC) SetUnicast() {
	m.SetIndividual()
}

// SetMulticast sets I/G bit of the MAC address to multicast (group)
func (m *MAC) SetMulticast() {
	m.SetGroup()
}

// SetIG sets if the address is individual via the I/G bit
func (m *MAC) SetIG(individual bool) {
	if individual {
		m.SetIndividual()
		return
	}
	m.SetGroup()
}

// SetOUI sets the OUI part of the MAC
func (m *MAC) SetOUI(oui [3]byte) {
	for i := 0; i < 3; i++ {
		m.b[i] = oui[i]
	}
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
