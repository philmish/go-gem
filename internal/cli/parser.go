package cli

import (
	"fmt"
	"github.com/philmish/go-gem/internal/config"
	"log"
	"os"
)

var envCommands = map[string]bool{
	"do":  true,
	"add": true,
	"ls":  true,
	"rm":  true,
}

func parseEnvCommand(u *UserInput) {
	if p, err := os.Getwd(); err == nil {
		project := config.FromFile(p)
		var env = project.Env

		switch u.Cmd {
		case "do":
			env.Do(u.Name, u.AddArgs)
		case "add":
			env.Add(u.Name, u.Arg, u.AddArgs)
			project.ToFile(env.WorkDir)
		case "ls":
			env.List()
		case "rm":
			env.Remove(u.Name)
			project.ToFile(env.WorkDir)
		}
	} else {
		log.Fatal("Failed to get current working directory")
	}
}

func (u *UserInput) Parse() {
	if ok, _ := envCommands[u.Cmd]; ok {
		parseEnvCommand(u)

	} else {
		switch u.Cmd {
		case "help":
			helpCmd()
		case "init":
			initParser(u)
		default:
			fmt.Printf("Unknown command: %s", u.Cmd)
		}
	}
}
