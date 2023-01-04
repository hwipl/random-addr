package cmd

import (
	"fmt"

	"github.com/hwipl/random-addr/internal/mac"
)

// printMAC prints m
func printMAC(m *mac.MAC) {
	fmt.Println("Random MAC Address")
	fmt.Println("==================")
	fmt.Println()
	fmt.Println(m)
	fmt.Println()
	fmt.Println("Details")
	fmt.Println("=======")
	fmt.Println()
	fmt.Println(m.Explain())
	fmt.Println()
	fmt.Println(m.Table())
	fmt.Println()
}

// Run is the main entry point
func Run() {
	m := mac.Random()
	printMAC(m)
}
