package templates

import (
        "github.com/philmish/go-gem/internal/environment"
)

func DefaultGoenv(workdir string) *environment.Environment {
        var nEnv = environment.NewEnv(workdir)

        nEnv.Add("test_all", "go test -v", []string{"./..."})
        
        return nEnv
}
