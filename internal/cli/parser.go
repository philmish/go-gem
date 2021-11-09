package cli

import (
        "fmt"
        "os"
        "log"
        "github.com/philmish/go-gem/internal/config"
)

type gemCmd struct {
        mod string
        arg string
        addargs []string
}

func (u *UserInput)Parse(){

        switch u.Cmd {
        case "do":
                var cmd = &gemCmd{u.Module, u.Arg, u.AddArgs}
                if p, err := os.Getwd(); err == nil {
                        project := config.FromFile(p)
                        NewShell(cmd, project)
                } else {
                        log.Fatalf("Failed to load current working directory.")
                }
        case "help":
                helpCmd()
        default:
                fmt.Printf("Unknown command: %s", u.Cmd)
        }
}
