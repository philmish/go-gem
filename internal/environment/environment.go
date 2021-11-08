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
        WorkDir string;
        SetupCmds map[string]Command `json:"setupcmds"`;
        BuildCmds map[string]Command `json:"buildcmds"`;
        RunCmds map[string]Command `json:"runcmds"`;
        TestCmds map[string]Command `json:"testcmds"`;
}

func NewEnv(dir string) *Environment {
        return &Environment{
                dir,
                make(map[string]Command),
                make(map[string]Command),
                make(map[string]Command),
                make(map[string]Command),
        }
}

func (e *Environment)getCmds(name string) map[string]Command {
        var cmds map[string]Command
        switch name {
        case "Setup":
                cmds = e.SetupCmds
        case "Build":
                cmds = e.BuildCmds
        case "Run":
                cmds = e.RunCmds
        case "Test":
                cmds = e.TestCmds
        default:
                log.Fatalf("%s is not a known command collection.", name)
        }
        return cmds
}

func (e *Environment)Add(module, name, cmd string, args []string){
        cmds := e.getCmds(module)
        if _, exists := cmds[name]; exists {
                log.Fatalf("%s already is a registered command.", name)
        } else {
                newCmd := Command{cmd, args}
                cmds[name] = newCmd
                fmt.Printf("%s was added to %s", name, module)
        }
}

func (e *Environment)Remove(module, name string) {
        cmds := e.getCmds(module)
        if _, exists := cmds[name]; exists {
                delete(cmds, name)
                fmt.Printf("%s removed from %s", name, module)
        } else {
                log.Fatalf("%s is not a command in %s", name, module)
        }
}

func (e *Environment)List(module string) {
        cmds := e.getCmds(module)
        for name, i := range cmds {
                fmt.Printf("Command Alias: %s\n", name)
                i.PrintCmd()
                fmt.Println("")
        }
}

func (e *Environment)Do(module, cmd string) {
        cmds := e.getCmds(module)
        if command, ok := cmds[cmd]; ok {
                err := command.Execute()

                if err != nil {
                        log.Fatalf("Execution of command: %s\n produced following error:\n%s", command, err.Error())
                }
        } else {
                log.Fatalf("%s is not a registered command.", cmd)
        }
}

type Setup interface {
        Add()
        Remove()
        List()
        Do()
}

