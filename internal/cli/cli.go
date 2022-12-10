package cli

import (
	"flag"
	"fmt"

	"github.com/philmish/go-gem/internal/config"
	"github.com/philmish/go-gem/internal/parser"
)

func GetInput() *parser.UserInput {
	var (
		cmd   string
		name  string
		arg   string
		alias bool
	)

	flag.StringVar(&cmd, "c", "help", "gem command you want to run.")
	flag.StringVar(&name, "n", "", "gem module to use for parsing the args.")
	flag.StringVar(&arg, "a", "", "argument passed to the gem command.")
	flag.BoolVar(&alias, "alias", false, "create an alias file.")
	flag.Parse()

	var data = &parser.UserInput{Cmd: cmd, Name: name, Arg: arg, Alias: alias}
	return data
}

func Parse(p *config.Project, u *parser.UserInput) {
	if ok, _ := parser.EnvCommands[u.Cmd]; ok {
		parser.ParseEnvCommand(p, u)

	} else {
		switch u.Cmd {
		case "help":
			helpCmdParser(u.Name)
		case "init":
			initParser(u)
		default:
			fmt.Printf("Unknown command: %s", u.Cmd)
		}
	}
}
