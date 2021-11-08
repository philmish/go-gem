package templates

import (
        "github.com/philmish/go-gem/internal/environment"
        "fmt"
)

func DefaultPyenv(workdir string) *environment.Environment{
        var nEnv = environment.NewEnv(workdir)
        var envPath = fmt.Sprintf("%s/venv/bin/activate", workdir)

        nEnv.Add("Setup", "env", "source", []string{envPath})
        nEnv.Add("Setup", "reqs", "pip install -r", []string{"requirements.txt"})
        nEnv.Add("Setup", "get", "pip install -U", []string{})

        return nEnv
}
