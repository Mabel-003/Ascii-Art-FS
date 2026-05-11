package alignment

import (
	"log"
	"os"
	"strings"
)

func PrintAsciiArt(text, banner string) string {
	content, err := os.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		log.Print(err.Error())
		return ""
	}
	arts := strings.Split(string(content), "\n")

	text = strings.ReplaceAll(text, "\r\n", "\n")
	words := strings.Split(text, "\n")
	result := strings.Builder{}

	for _, word := range words {
		if word == "" {
			result.WriteByte('\n')
			continue
		}
		for row := 0; row < 8; row++ {
			for _, ch := range word {
				index := (int(ch) - 32) * 9
				result.WriteString(arts[index+row+1])
			}
			result.WriteByte('\n')
		}
	}
	return result.String()

}
