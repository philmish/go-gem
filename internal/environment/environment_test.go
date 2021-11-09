package environment_test

import (
        "testing"
        "github.com/philmish/go-gem/internal/environment"
)

func TestCommand(t *testing.T) {
        newCmd := environment.Command{"echo",[]string{"Hello World"},}
        newCmd.PrintCmd()
        newCmd.Execute()
}

func TestEnvironment(t *testing.T) {
        var newEnv = environment.NewEnv("/foo/bar")
        newEnv.Add("hello", "echo", []string{"Hello World"})
        newEnv.List()
        newEnv.Do("hello", []string{"with", "args"})
        newEnv.Remove("hello")
        newEnv.List()
}
