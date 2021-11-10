package main

import (
	"github.com/philmish/go-gem/internal/cli"
)

func main() {
	command := cli.GetInput()
	command.Parse()
}
