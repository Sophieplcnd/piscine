package main

import (
	"fmt"
	"os"
)

func main() {
	threeStr := os.Args[1:]
	for _, v := range threeStr {
		if (v == "01") || (v == "galaxy") || (v == "galaxy 01") {
			fmt.Println("Alert!!!")
			break
		}
	}
}
