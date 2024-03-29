package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	myProgramName := os.Args[0]
	s := []rune(myProgramName)
	for index, name := range s {
		// if index not > 1, the print will include ./, the index 0 will be ., and the index 1 will be /
		if index > 1 {
			z01.PrintRune(name)
		}
	}
	z01.PrintRune('\n')
}
