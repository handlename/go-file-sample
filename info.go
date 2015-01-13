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
	// Lstat
	lstat, err := os.Lstat("sample.txt")
	if err != nil {
		fmt.Println("error on Lstat:", err)
		return false
	}

	pp.Println(lstat)

	// Stat
	stat, err := os.Stat("sample.txt")
	if err != nil {
		fmt.Println("error on Stat:", err)
		return false
	}

	pp.Println(stat)

	return true
}
