package cmd

import (
	"fmt"

	"github.com/hwipl/random-addr/internal/mac"
)

// Run is the main entry point
func Run() {
	fmt.Println(mac.Random())
}
