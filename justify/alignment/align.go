package alignment

import (
	"strings"
	"golang.org/x/term"
	"os"
)

// Get terminal width safely
func getTerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 80
	}
	return width
}

// -------- LEFT ALIGN --------
func AlignLeft(ascii string) string {
	return ascii // already left aligned
}

// -------- RIGHT ALIGN --------
func AlignRight(ascii string) string {
	lines := strings.Split(ascii, "\n")
	width := getTerminalWidth()

	var result strings.Builder

	for _, line := range lines {
		if line == "" {
			result.WriteByte('\n')
			continue
		}

		pad := width - len(line)
		if pad < 0 {
			pad = 0
		}

		result.WriteString(strings.Repeat(" ", pad))
		result.WriteString(line)
		result.WriteByte('\n')
	}

	return result.String()
}

// -------- CENTER ALIGN --------
func AlignCenter(ascii string) string {
	lines := strings.Split(ascii, "\n")
	width := getTerminalWidth()

	var result strings.Builder

	for _, line := range lines {
		if line == "" {
			result.WriteByte('\n')
			continue
		}

		pad := (width - len(line)) / 2
		if pad < 0 {
			pad = 0
		}

		result.WriteString(strings.Repeat(" ", pad))
		result.WriteString(line)
		result.WriteByte('\n')
	}

	return result.String()
}

// -------- JUSTIFY ALIGN --------
func AlignJustify(text, banner string) string {
	// IMPORTANT: we must rebuild word-by-word ASCII
	content, _ := os.ReadFile("banners/" + banner + ".txt")
	arts := strings.Split(string(content), "\n")

	width := getTerminalWidth()
	lines := strings.Split(text, "\n")

	var result strings.Builder

	for _, line := range lines {
		if line == "" {
			result.WriteByte('\n')
			continue
		}

		words := strings.Fields(line)

		// fallback to normal ASCII if only one word
		if len(words) <= 1 {
			result.WriteString(PrintAsciiArt(line, banner))
			continue
		}

		// build ascii blocks for each word
		wordBlocks := make([][]string, len(words))

		for i, word := range words {
			block := make([]string, 8)

			for row := 0; row < 8; row++ {
				for _, ch := range word {
					index := (int(ch) - 32) * 9
					block[row] += arts[index+row+1]
				}
			}
			wordBlocks[i] = block
		}

		// calculate width of words (without spacing)
		wordWidth := 0
		for _, w := range wordBlocks {
			wordWidth += len(w[0])
		}

		extraSpace := (width - wordWidth) / (len(words) - 1)
		if extraSpace < 0 {
			extraSpace = 0
		}

		// print row by row
		for row := 0; row < 8; row++ {
			for i, block := range wordBlocks {
				result.WriteString(block[row])
				if i != len(wordBlocks)-1 {
					result.WriteString(strings.Repeat(" ", extraSpace))
				}
			}
			result.WriteByte('\n')
		}
	}

	return result.String()
}