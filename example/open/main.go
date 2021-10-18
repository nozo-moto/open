package main

import (
	"os"

	"github.com/nozo-moto/open"
)

func main() {
	args := os.Args
	if len(args) < 1 {
		panic("set argument")
	}
	if err := open.Open(args[1:]...); err != nil {
		panic(err)
	}
}
