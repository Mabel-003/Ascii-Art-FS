package main

import (
	"bytes"
	"os"
	"testing"
)

func TestAsciiArtOutput(t *testing.T) {

	// Save original stdout
	oldStdout := os.Stdout

	// Create a pipe to capture output
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Fake command line argument
	os.Args = []string{"cmd", "A"}

	// Run main function
	main()

	// Close writer
	w.Close()

	// Restore stdout
	os.Stdout = oldStdout

	// Read captured output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Check if output is empty
	if output == "" {
		t.Errorf("Expected ASCII art output but got empty output")
	}
}
