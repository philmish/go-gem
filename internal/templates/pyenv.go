package templates

import (
	"fmt"
	"github.com/philmish/go-gem/internal/environment"
)

func DefaultPyenv(workdir string, aliasing bool) *environment.Environment {
	var nEnv = environment.NewEnv(workdir)
	nEnv.Alias = aliasing
	var envPath = fmt.Sprintf("%s/venv/bin/activate", workdir)

	nEnv.Add("menv", "python3", []string{"-m", "venv", "venv"})
	nEnv.Add("senv", "source", []string{envPath})
	nEnv.Add("reqs", "pip", []string{"install", "-r", "requirements.txt"})
	nEnv.Add("get", "pip", []string{"install", "-U"})

	return nEnv
}
