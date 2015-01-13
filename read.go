package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if !doMain() {
		os.Exit(1)
	}
}

func doMain() bool {
	path := flag.String("path", "", "target file path")
	flag.Parse()

	file, err := os.Open(*path)
	if err != nil {
		fmt.Println("error on Open:", err)
		return false
	}
	defer file.Close()

	fmt.Println(file.Name())

	var buf []byte
	n, err := file.Read(buf)
	if err != nil {
		fmt.Println("error on Read:", err)
		return false
	}

	fmt.Println("n =", n)
	fmt.Println(buf)

	return true
}
