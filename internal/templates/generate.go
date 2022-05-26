package templates

import (
	"errors"
	"github.com/philmish/go-gem/internal/environment"
	"log"
)

type aliasTemplate struct {
	name string
	cmd  string
	args []string
}

func addTemplateCmd(e *environment.Environment, template aliasTemplate) {
	e.Add(template.name, template.cmd, template.args)
}

func addTemplateList(e *environment.Environment, templates []aliasTemplate) {
	for _, template := range templates {
		addTemplateCmd(e, template)
	}
}

func generate(workdir string, aliasing bool, templates []aliasTemplate) *environment.Environment {
	var newEnv = environment.NewEnv(workdir)
	newEnv.Alias = aliasing
	defs := append(gogemDefaults, gitDefaults...)
	templates = append(templates, defs...)
	addTemplateList(newEnv, templates)
	return newEnv
}

var defaultMap = map[string][]aliasTemplate{
	"node": {
		aliasTemplate{"v", "nvm", []string{"use"}},
		aliasTemplate{"vinst", "nvm", []string{"install"}},
		aliasTemplate{"get", "npm", []string{"install"}},
		aliasTemplate{"run", "npm", []string{"install"}},
	},
	"python": {
		aliasTemplate{"menv", "python3", []string{"-m", "venv", "venv"}},
		aliasTemplate{"reqs", "pip", []string{"install", "-r", "requirements.txt"}},
		aliasTemplate{"get", "pip", []string{"install", "-U"}},
	},
	"vue": {
		aliasTemplate{"get", "npm", []string{"install"}},
		aliasTemplate{"dev", "npm", []string{"run", "serve"}},
	},
	"go": {
		aliasTemplate{"test_all", "go", []string{"test", "-v", "./..."}},
		aliasTemplate{"tidy", "go", []string{"mod", "tidy"}},
		aliasTemplate{"fmt", "go", []string{"fmt", "./..."}},
	},
}

var gogemDefaults = []aliasTemplate{
	{"gemls", "gogem", []string{"-c", "ls"}},
	{"gemtodo", "gogem", []string{"-c", "lstodo"}},
	{"gemdone", "gogem", []string{"-c", "lsdone"}},
	{"gemdo", "gogem", []string{"-c", "do", "-n"}},
	{"gemadd", "gogem", []string{"-c", "add", "-n"}},
	{"gemrm", "gogem", []string{"-c", "rm", "-n"}},
}

var gitDefaults = []aliasTemplate{
	{"ga", "git", []string{"add", "."}},
	{"gs", "git", []string{"status"}},
	{"gc", "git", []string{"commit", "-m"}},
	{"push", "git", []string{"push"}},
}

func CreateTemplate(name, workdir string, aliasing bool) (*environment.Environment, error) {
	if template, ok := defaultMap[name]; ok {
		return generate(workdir, aliasing, template), nil
	} else {
		log.Println("CreateTemplate failed")
		return nil, errors.New("Name does not exist")
	}
}

func CreateDefault(workdir string, aliasing bool) *environment.Environment {
	temps := []aliasTemplate{}
	return generate(workdir, aliasing, temps)

}
