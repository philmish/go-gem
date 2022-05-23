package templates

import (
	"github.com/philmish/go-gem/internal/environment"
)

func DefaultVueEnv(workdir string, aliasing bool) *environment.Environment {
	var newEnv = environment.NewEnv(workdir)
	newEnv.Alias = aliasing

	newEnv.Add("get", "npm", []string{"install"})
	newEnv.Add("dev", "npm", []string{"run", "serve"})

	return newEnv
}
