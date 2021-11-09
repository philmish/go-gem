package templates

import (
        "github.com/philmish/go-gem/internal/environment"
        "fmt"
)

func DefaultPyenv(workdir string) *environment.Environment{
        var nEnv = environment.NewEnv(workdir)
        var envPath = fmt.Sprintf("%s/venv/bin/activate", workdir)

        nEnv.Add("env", "source", []string{envPath})
        nEnv.Add("reqs", "pip install -r", []string{"requirements.txt"})
        nEnv.Add("get", "pip install -U", []string{})

        return nEnv
}
