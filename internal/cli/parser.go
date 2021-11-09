package cli

import (
        "fmt"
        "os"
        "log"
        "github.com/philmish/go-gem/internal/config"
)

var envCommands = map[string]bool{
        "do": true,
        "add": true,
        "ls": true,
        "rm": true,
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
                case "ls":
                        env.List()
                case "rm":
                        env.Remove(u.Name)
                }
        } else {
                log.Fatal("Failed to get current working directory")
        }
}

func (u *UserInput)Parse(){
        if ok, _ := envCommands[u.Cmd]; ok {
                parseEnvCommand(u)

        } else {
                switch u.Cmd {
                case "help":
                        helpCmd()
                default:
                        fmt.Printf("Unknown command: %s", u.Cmd)
                }
        }
}
