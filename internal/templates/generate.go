package templates

import (
	"errors"
	"fmt"

	"github.com/philmish/go-gem/internal/environment"
)

type aliasTemplate struct {
	name string
	cmd  string
	args []string
}

func generate(workdir string, aliasing bool, templates []*aliasTemplate) *environment.Environment {
	var newEnv = environment.NewEnv(workdir)
	newEnv.Alias = aliasing

	for _, template := range templates {
		newEnv.Add(template.name, template.cmd, template.args)
	}
	return newEnv
}

var defaultMap = map[string][]aliasTemplate{
	"node": []aliasTemplate{
		aliasTemplate{"v", "nvm", []string{"use"}},
		aliasTemplate{"vinst", "nvm", []string{"install"}},
		aliasTemplate{"get", "npm", []string{"install"}},
		aliasTemplate{"run", "npm", []string{"install"}},
	},
	"python": []aliasTemplate{
		aliasTemplate{"menv", "python3", []string{"-m", "venv", "venv"}},
		aliasTemplate{"reqs", "pip", []string{"install", "-r", "requirements.txt"}},
		aliasTemplate{"get", "pip", []string{"install", "-U"}},
	},
	"vue": []aliasTemplate{
		aliasTemplate{"get", "npm", []string{"install"}},
		aliasTemplate{"dev", "npm", []string{"run", "serve"}},
	},
	"go": []aliasTemplate{
		aliasTemplate{"test_all", "go", []string{"test", "-v", "./..."}},
		aliasTemplate{"tidy", "go", []string{"mod", "tidy"}},
		aliasTemplate{"fmt", "go", []string{"fmt", "./..."}},
	},
}

func CreateTemplate(name, workdir string, aliasing bool) (*environment.Environment, error) {
	if template, ok := defaultMap[name]; ok {
		return generate(workdir, aliasing, &template), nil
	} else {
		return nil, errors.New("Name does not exist")
	}
}
