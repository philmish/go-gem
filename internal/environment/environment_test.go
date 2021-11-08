package environment_test

import (
        "testing"
        "github.com/philmish/go-gem/internal/environment"
)

func TestCommand(t *testing.T) {
        newCmd := Command{"echo",[]string{"Hello World"},}
        newCmd.print_cmd()
        newCmd.Execute()
}
