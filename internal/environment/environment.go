package environment

import (
        "log"
        "os"
        "os/exec"
        "fmt"
        "strings"
)

type Command struct {
        Name string `json:"name"`;
        Args []string `json:"args"`;
}

func (c *Command)PrintCmd() {
        data := fmt.Sprintf("Name: %s | Args: %s", c.Name, strings.Join(c.Args[:], ","))
        fmt.Println(data)
}

func (c *Command)Execute() error {
        cmd := exec.Command(c.Name, c.Args...)
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        return cmd.Run()
}

func (c *Command)updateArgs(newArgs []string) {
        c.Args = newArgs
}

type Environment struct {
        WorkDir string `json:"workdir"`
        Cmds map[string]Command `json:"cmds"`
}

func NewEnv(dir string) *Environment {
        return &Environment{
                dir,
                make(map[string]Command),
        }
}

func (e *Environment)Add(name, cmd string, args []string){
        if _, exists := e.Cmds[name]; exists {
                log.Fatalf("%s already is a registered command.", name)
        } else {
                newCmd := Command{cmd, args}
                e.Cmds[name] = newCmd
                fmt.Printf("%s was added.", name)
        }
}

func (e *Environment)Remove(name string) {
        if _, exists := e.Cmds[name]; exists {
                delete(e.Cmds, name)
                fmt.Printf("%s was removed.", name)
        } else {
                log.Fatalf("%s is not a command.", name)
        }
}

func (e *Environment)List() {
        for name, i := range e.Cmds {
                fmt.Printf("Command Alias: %s\n", name)
                i.PrintCmd()
                fmt.Println("")
        }
}

func (e *Environment)Do(name string, addargs []string) {
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

