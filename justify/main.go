package main

import (
	"flag"
	"fmt"
	"ascii-art-justify/alignment"
)

func main() {
	alignType := flag.String("align", "left", "alignment type: left, center, right, justify")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: go run . [--align=<type>] <text> [<banner>]")
		return
	}

	text := args[0]
	banner := "standard"
	if len(args) >= 2 {
		banner = args[1]
	}

	ascii := alignment.PrintAsciiArt(text, banner)
	ascii = alignment.SelectAlign(*alignType, ascii, text, banner)

	fmt.Print(ascii)
}
