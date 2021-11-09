package templates

import (
        "github.com/philmish/go-gem/internal/environment"
        "fmt"
)

func DefaultPyenv(workdir string) *environment.Environment{
        var nEnv = environment.NewEnv(workdir)
        var envPath = fmt.Sprintf("%s/venv/bin/activate", workdir)

        nEnv.Add("env", "source", []string{envPath})
        nEnv.Add("reqs", "pip", []string{"install", "-r", "requirements.txt"})
        nEnv.Add("get", "pip", []string{"install", "-U"})

        return nEnv
}
