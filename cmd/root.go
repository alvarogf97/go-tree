package cmd

import (
	"fmt"
	"os"
)

// Cobra entry point
func Execute() {
	if err := TreeCommand().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
