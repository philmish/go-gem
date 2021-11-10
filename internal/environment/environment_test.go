package environment_test

import (
        "testing"
        "os"
        "github.com/philmish/go-gem/internal/environment"
)

func checkErr(e error, t *testing.T) {
        if e != nil {
                t.Errorf("%v", e)
        }
}

func TestCommand(t *testing.T) {
        newCmd := environment.Command{"echo",[]string{"Hello World"},}
        newCmd.PrintCmd()
        newCmd.Execute()
}

func TestEnvironment(t *testing.T) {
        p, err := os.Getwd()
        checkErr(err, t)
        var newEnv = environment.NewEnv(p)

        newEnv.Add("hello", "echo", []string{"Hello World"})
        newEnv.List()
        newEnv.Do("hello", []string{"with", "args"})
        err = newEnv.Aliases()
        checkErr(err, t)
        newEnv.Remove("hello")
        newEnv.List()
}
