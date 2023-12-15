package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	// Save original command-line arguments
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Set test command-line arguments
	os.Args = []string{"./chrnr-cli", "arg1", "arg2"}

	// Call the main function
	main()
}
