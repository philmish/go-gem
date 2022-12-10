package main

import (
	"fmt"
	"os"

	"github.com/philmish/go-gem/internal/cli"
	"github.com/philmish/go-gem/internal/config"
	"github.com/philmish/go-gem/internal/shell"
)

func main() {
    if ok, _ := config.WorkDirHasConf(); !ok {
        fmt.Println("Could not find config file.")
        os.Exit(1)
    } 
    conf := config.FromFile("./")
    if os.Args[1] == "shell" {
        shell.GemShell{Project: *conf}.RunShell()
    } else {
        command := cli.GetInput()
        cli.Parse(conf, command)
    }
}
