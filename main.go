package main

import (
	"fmt"
	"os"

	"github.com/levpaul/golarity/src/magnet"
)

func main() {
	var magnetURL string
	if len(os.Args) == 2 {
		magnetURL = os.Args[1]
	} else {
		fmt.Println("FATAL: expected a magnet to be passed as an arg")
		os.Exit(1)
	}

	mag, err := magnet.ParseMagnet(magnetURL)
	if err != nil {
		fmt.Printf("Found error: '%w' - exiting", err)
		os.Exit(1)
	}
	mag.PrintInfo()
}
