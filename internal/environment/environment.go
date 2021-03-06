package environment

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Environment struct {
	WorkDir string             `json:"workdir"`
	Cmds    map[string]Command `json:"cmds"`
	Alias   bool               `json:"alias"`
}

func NewEnv(dir string) *Environment {
	return &Environment{
		dir,
		make(map[string]Command),
		false,
	}
}

func (e *Environment) Add(name, cmd string, args []string) {
	if _, exists := e.Cmds[name]; exists {
		log.Fatalf("%s already is a registered command.", name)
	} else {
		newCmd := Command{cmd, args}
		e.Cmds[name] = newCmd
		fmt.Printf("%s was added.\n", name)
	}
}

func (e *Environment) Remove(name string) {
	if _, exists := e.Cmds[name]; exists {
		delete(e.Cmds, name)
		fmt.Printf("%s was removed.", name)
	} else {
		log.Fatalf("%s is not a command.", name)
	}
}

func (e *Environment) List() {
	for name, i := range e.Cmds {
		fmt.Printf("Command Alias: %s\n", name)
		i.PrintCmd()
		fmt.Println("")
	}
}

func (e *Environment) Do(name string, addargs []string) {
	if command, ok := e.Cmds[name]; ok {
		for _, i := range addargs {
			command.Args = append(command.Args, i)
		}
		err := command.Execute()

		if err != nil {
			log.Fatalf("Execution of command: %s\n produced following error:\n%s", command, err.Error())
		}
	} else {
		log.Fatalf("%s is not a registered command.", name)
	}
}

func (e *Environment) Aliases() error {
	var aliases []string
	for k, v := range e.Cmds {
		aliases = append(aliases, v.alias(k))
	}
	inp := strings.Join(aliases[:], "\n")
	data := []byte(inp)
	fpath := fmt.Sprintf("%s/.gem_aliases", e.WorkDir)
	return os.WriteFile(fpath, data, 0644)
}
