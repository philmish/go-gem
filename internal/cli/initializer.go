package cli

import (
	"github.com/philmish/go-gem/internal/config"
	"github.com/philmish/go-gem/internal/environment"
	"github.com/philmish/go-gem/internal/templates"
	"log"
	"os"
)

func newProject(name string, e *environment.Environment) *config.Project {
	newPr := config.NewProject()
	newPr.Name = name
	newPr.Env = *e
	return newPr

}

func defaultEnv(name string, aliasing bool) {

	if p, err := os.Getwd(); err == nil {
		newEnv := environment.NewEnv(p)
		newEnv.Alias = aliasing
		np := newProject(name, newEnv)
		np.ToFile(p)
		log.Printf("%s created successfully.\n", name)
	} else {
		log.Fatal("Something went wrong when trying to read current working directory.")
	}
}

func writeEnv(name string, env *environment.Environment) {
	np := newProject(name, env)
	p := env.WorkDir
	np.ToFile(p)
	log.Printf("%s created successfully.", name)
}

func createEnv(envtype string, name string, aliasing bool) {
	p, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to load working directory\nError:\n%v", err)
	}
	switch envtype {
	case "go":
		newEnv := templates.DefaultGoenv(p, aliasing)
		writeEnv(name, newEnv)
	case "py":
		newEnv := templates.DefaultPyenv(p, aliasing)
		writeEnv(name, newEnv)
	case "node":
		newEnv := templates.DefaultNodeEnv(p, aliasing)
		writeEnv(name, newEnv)
	default:
		defaultEnv(name, aliasing)
	}
}

func initParser(i *UserInput) {
	createEnv(i.Name, i.Arg, i.Alias)
}
