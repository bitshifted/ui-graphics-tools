package main

import (
	"log"
	"os"

	"github.com/bitshifted/ui-graphics-tools/cli"
)

const (
	exitError = 10
)

func main() {
	err := cli.Run()
	if err != nil {
		log.Println(err.Error())
		os.Exit(exitError)
	}
}
