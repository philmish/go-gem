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
