package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: parsefile filename.dsv")
		os.Exit(1)
	}

	parseGameFile(args[1])
}
