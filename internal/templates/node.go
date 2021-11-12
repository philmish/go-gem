package templates

import (
	"github.com/philmish/go-gem/internal/environment"
)

func DefaultNodeEnv(workdir string, aliasing bool) *environment.Environment {
	var newEnv = environment.NewEnv(workdir)
	newEnv.Alias = aliasing

	newEnv.Add("v", "nvm", []string{"use"})
	newEnv.Add("vinst", "nvm", []string{"install"})
	newEnv.Add("get", "npm", []string{"install"})

	return newEnv
}
