package main

import (
	"github.com/bryanaustin/version"
)

func main() {
	c := version.ConfigureFromArgs()
	version.Process(c)
}
