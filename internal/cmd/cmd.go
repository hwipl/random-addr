package cmd

import (
	"flag"
	"fmt"

	"github.com/hwipl/random-addr/internal/ipv4"
	"github.com/hwipl/random-addr/internal/ipv6"
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

// runMAC runs the mac subcommand
func runMAC() {
	m := mac.Random()
	printMAC(m)
}

// runIPv4 runs the ipv4 subcommand
func runIPv4() {
	ip := ipv4.Random()
	fmt.Println(ip)
}

// runIPv6 runs the ipv6 subcommand
func runIPv6() {
	ip := ipv6.Random()
	fmt.Println(ip)
}

// Run is the main entry point
func Run() {
	flag.Parse()
	switch flag.Arg(0) {
	case "mac":
		runMAC()
	case "ipv4":
		runIPv4()
	case "ipv6":
		runIPv6()
	default:
		runMAC()
	}
}
