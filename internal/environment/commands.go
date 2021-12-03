package environment

import (
        "fmt"
        "os"
        "os/exec"
        "strings"
)


type Command struct {
	Name string   `json:"name"`
	Args []string `json:"args"`
}

func (c *Command) PrintCmd() {
	data := fmt.Sprintf("Command: %s | Args: %s", c.Name, strings.Join(c.Args[:], ","))
	fmt.Println(data)
}

func (c *Command) alias(name string) string {
	return fmt.Sprintf("alias %s=\"%s %s\"", name, c.Name, strings.Join(c.Args[:], " "))
}

func (c *Command) Execute() error {
	cmd := exec.Command(c.Name, c.Args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (c *Command) updateArgs(newArgs []string) {
	c.Args = newArgs
}
