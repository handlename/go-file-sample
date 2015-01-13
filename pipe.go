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
	// 内部でsyscall.Pipeを呼んでいる
	f1, f2, err := os.Pipe()
	if err != nil {
		fmt.Println("error on Pipe:", err)
		return false
	}
	defer f1.Close()
	defer f2.Close()

	pp.Println(f1)
	pp.Println(f2)

	_, err = f2.WriteString("hello")
	if err != nil {
		fmt.Println("error on WriteString:", err)
		return false
	}

	return true
}
