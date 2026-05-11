package alignment

import (
	"os"
	"strings"
	"testing"
)

func withParentDir(t *testing.T, fn func()) {
    oldDir, err := os.Getwd()
    if err != nil {
        t.Fatal(err)
    }
    if err := os.Chdir("../"); err != nil {
        t.Fatal(err)
    }
    defer os.Chdir(oldDir)
    fn()
}

func TestSelectAlignLeft(t *testing.T) {
	text := "hello"
	banner := "standard"
	ascii := PrintAsciiArt(text, banner)
	result := SelectAlign("left", ascii, text, banner)
	if result != ascii {
		t.Errorf("Left align should return the original ascii art")
	}
}

func TestSelectAlignCenter(t *testing.T) {
	withParentDir(t, func() {
	text := "hi"
	banner := "shadow"
	ascii := PrintAsciiArt(text, banner)
	result := SelectAlign("center", ascii, text, banner)
	lines := strings.Split(result, "\n")
	hasLeadingSpace := false
	for _, line := range lines {
		if strings.HasPrefix(line, " ") {
			hasLeadingSpace = true
			break
		}
	}
	if !hasLeadingSpace {
		t.Errorf("Center align should have at least one line with leading spaces")
	}
	})
}

func TestSelectAlignRight(t *testing.T) {
	withParentDir(t, func() {
	text := "hi"
	banner := "standard"
	ascii := PrintAsciiArt(text, banner)
	result := SelectAlign("right", ascii, text, banner)
	lines := strings.Split(result, "\n")
	hasSpace := false
	for _, line := range lines {
		if strings.Contains(line, " ") {
			hasSpace = true
			break
		}
	}
	if !hasSpace {
		t.Errorf("Right align should have spaces in lines")
	}
	})
}

func TestSelectAlignJustify(t *testing.T) {
	withParentDir(t, func() {
	text := "hello world"
	banner := "thinkertoy"
	ascii := PrintAsciiArt(text, banner)
	result := SelectAlign("justify", ascii, text, banner)
	if result == "" {
		t.Errorf("Justify should produce output")
	}
	// For justify with multiple words, it should have spaces between words
	if !strings.Contains(result, " ") {
		t.Errorf("Justify should distribute spaces")
	}
	})
}
