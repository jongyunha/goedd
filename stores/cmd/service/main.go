package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("stores exited abnormally: %s\n", err)
		os.Exit(1)
	}
}

func run() (err error) {
	// TODO
	return
}
