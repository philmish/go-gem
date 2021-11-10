package cli

import (
	"github.com/philmish/go-gem/internal/config"
	"github.com/philmish/go-gem/internal/environment"
	"log"
	"os"
)

func newProject(name string, e *environment.Environment) *config.Project {
	newPr := config.NewProject()
	newPr.Name = name
	newPr.Env = *e
	return newPr

}

func EmptyDefault(name string, aliasing bool) {

	if p, err := os.Getwd(); err == nil {
		newEnv := environment.NewEnv(p)
		newEnv.Alias = aliasing
		np := newProject(name, newEnv)
		np.ToFile(p)
		log.Printf("%s created successfully.", name)
	} else {
		log.Fatal("Something went wrong when trying to read current working directory.")
	}
}

func initParser(i *UserInput) {
	switch i.Name {
	case "empty":
		EmptyDefault(i.Arg, i.Alias)
	default:
		EmptyDefault("default", i.Alias)
	}
}
