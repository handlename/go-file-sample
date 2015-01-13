package main

import (
	"fmt"
	"os"

	"github.com/k0kubun/pp"
)

func main() {
	if !doMain() {
		os.Exit(1)
	}
}

func doMain() bool {
	file, err := os.Open(".")
	if err != nil {
		fmt.Println("error on Open:", err)
		return false
	}
	defer file.Close()

	// Readdir
	infoList, err := file.Readdir(0)
	if err != nil {
		fmt.Println("error on Readdir:", err)
		return false
	}

	pp.Println(infoList)

	// Seek
	if _, err := file.Seek(0, 0); err != nil {
		fmt.Println("error on Seek:", err)
		return false
	}

	// Readdirnames
	names, err := file.Readdirnames(0)
	if err != nil {
		fmt.Println("error on Readdirnames:", err)
		return false
	}

	pp.Println(names)

	return true
}
