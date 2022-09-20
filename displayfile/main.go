package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	argument := os.Args
	if len(argument) == 1 {
		fmt.Println("File name missing")
		return
	}
	if len(argument) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	bs, err := ioutil.ReadFile(argument[1])
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Print(str)
}
