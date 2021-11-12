package templates

import (
	"github.com/philmish/go-gem/internal/environment"
)

func DefaultGoenv(workdir string, aliasing bool) *environment.Environment {
	var nEnv = environment.NewEnv(workdir)
	nEnv.Alias = aliasing
	nEnv.Add("test_all", "go", []string{"test", "-v", "./..."})
	nEnv.Add("tidy", "go", []string{"mod", "tidy"})
	nEnv.Add("fmt", "go", []string{"fmt", "./..."})

	return nEnv
}
