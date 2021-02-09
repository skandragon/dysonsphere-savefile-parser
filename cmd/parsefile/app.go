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

func checkVers(b *Buffer, expected int32, name string) {
	vers := b.GetInt32le()
	fmt.Printf("Parsing %s, vers %d...\n", name, vers)
	if vers != expected {
		panic(fmt.Sprintf("Unknown %s version: %d, expected %d", name, vers, expected))
	}
}
