package cli

import (
        "fmt"
        "github.com/philmish/go-gem/internal/config"
)

var shellMods = map[string]bool{
        "do": true,
        "add": true,
        "ls": true,
        "rm": true,
}



func NewShell(cmd string, c *gemCmd, p *config.Project) {
        if ok, err := shellMods[cmd]; ok {
                switch cmd {
                case "do":
                        p.Env.Do(c.mod, c.arg, c.addargs)
                case "add":
                        p.Env.Add(c.mod, c.arg, c.addargs)
                }
        }

        fmt.Println("Implement this")
}


