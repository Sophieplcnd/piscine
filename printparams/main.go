package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args1 := os.Args
	for index, value := range args1 {
		if index != 0 {
			for _, value := range value {
				z01.PrintRune(value)
			}
			z01.PrintRune('\n')
		}
	}
}
