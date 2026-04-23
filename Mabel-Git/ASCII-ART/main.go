package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: wrong number of arguments")
		return
	}

	input := os.Args[1]

	if input == "" {
		os.Exit(0)
		
	}

	if strings.ReplaceAll(input, "\\n", "") == "" {
		count := len(input) / 2
		fmt.Print(strings.Repeat("\n", count))
		return
	}

	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error: file not in directory")
		return
	}
	text := string(data)
	lines := strings.Split(text, "\n")
	userInput := strings.Split(input, "\\n")
	for _, word := range userInput {
		if word != "" {
			for row := 0; row < 8; row++ {
				for _, char := range word {
					indexCalc := (int(char)-32)*9 + 1 + row
					fmt.Print(lines[indexCalc])
				}
				fmt.Println()
			}
		} else {
			fmt.Println()
		}
	}

}
