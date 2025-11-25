package main

import "github.com/01-edu/z01"

func main() {
	for j := 'z'; j >= 'a'; j-- {
		z01.PrintRune(j)
	}
	z01.PrintRune('\n')
}
