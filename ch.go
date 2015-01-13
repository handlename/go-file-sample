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

	file, err := os.Create(*path)
	if err != nil {
		fmt.Println("error on Create:", err)
		return false
	}
	defer file.Close()

	// Chmod
	err = file.Chmod(0777)
	if err != nil {
		fmt.Println("error on Chmod:", err)
		return false
	}

	// Chown
	err = file.Chown(501, -2)
	if err != nil {
		fmt.Println("error on Chown:", err)
		return false
	}

	// Sync
	err = file.Sync()
	if err != nil {
		fmt.Println("error on Sync:", err)
		return false
	}

	return true
}
