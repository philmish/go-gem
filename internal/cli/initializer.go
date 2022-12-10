package cli

import (
	"log"
	"os"

	"github.com/philmish/go-gem/internal/config"
	"github.com/philmish/go-gem/internal/environment"
	"github.com/philmish/go-gem/internal/parser"
	"github.com/philmish/go-gem/internal/templates"
)

func newProject(name string, e *environment.Environment) *config.Project {
	newPr := config.NewProject()
	newPr.Name = name
	newPr.Env = *e
	return newPr

}

func defaultProject(name string, aliasing bool) {

	if p, err := os.Getwd(); err == nil {
		newEnv := templates.CreateDefault(p, aliasing)
		np := newProject(name, newEnv)
		np.ToFile(p)
		log.Printf("%s created successfully.\n", name)
	} else {
		log.Fatal("Failed to read current working directory.")
	}
}

func writeEnv(name string, env *environment.Environment) {
	np := newProject(name, env)
	p := env.WorkDir
	np.ToFile(p)
	log.Printf("%s created successfully.", name)
}

func createProject(envtype, name string, aliasing bool) {
	p, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to load working directory\nError:\n%v", err)
	}
	env, err := templates.CreateTemplate(envtype, p, aliasing)
	if err != nil {
		log.Printf("Can't find template for %s\n. Creating default ...\n", envtype)
		defaultProject(name, aliasing)
	} else {
		writeEnv(name, env)
	}
}

func initParser(i *parser.UserInput) {
	createProject(i.Name, i.Arg, i.Alias)
}
