package parser

import (
    "log"
    "os"
    "fmt"
    "strconv"
    "github.com/philmish/go-gem/internal/config"
    
)

var EnvCommands = map[string]bool{
	"do":      true,
	"add":     true,
	"ls":      true,
	"rm":      true,
	"lstodo":  true,
	"lsdone":  true,
	"addtodo": true,
	"deltodo": true,
	"churg":   true,
    "shell":   true,
}

func ParseEnvCommand(project *config.Project, u *UserInput) {
	if _, err := os.Getwd(); err == nil {
		var env = &project.Env

		switch u.Cmd {
		case "do":
			env.Do(u.Name, u.AddArgs)
		case "add":
			env.Add(u.Name, u.Arg, u.AddArgs)
			project.ToFile(env.WorkDir)
		case "ls":
			env.List()
		case "rm":
			env.Remove(u.Name)
			project.ToFile(env.WorkDir)
		case "lstodo":
			todos := project.ListTodos(false)
			fmt.Printf("%s", todos)
		case "lsdone":
			todos := project.ListTodos(true)
			fmt.Printf("%s", todos)
		case "addtodo":
			urg, err := strconv.Atoi(u.Arg)
			if err != nil {
				log.Fatalf("Could not convert urgency %s to an number\n%v", u.Arg, err)
			}
			urgency := int8(urg)
			err = project.AddTodo(u.Name, urgency)
			if err != nil {
				log.Fatalf("Could not add todo.\n%v", err)
			}
			project.ToFile(env.WorkDir)
		case "deltodo":
			id, err := strconv.Atoi(u.Name)
			if err != nil {
				log.Fatalf("Could not convert id %s to an number\n%v", u.Name, err)
			}
			err = project.DelTodo(id)
			if err != nil {
				log.Fatalf("Could not delete todo with id %d\n%v", id, err)
			}
			project.ToFile(env.WorkDir)
		case "churg":
			id, err := strconv.Atoi(u.Name)
			if err != nil {
				log.Fatalf("Could not convert id %s to an number\n%v", u.Name, err)
			}
			urg, err := strconv.Atoi(u.Arg)
			if err != nil {
				log.Fatalf("Could not convert urgency %s to an number\n%v", u.Arg, err)
			}
			urgency := int8(urg)
			err = project.ChangeUrgency(id, urgency)
			if err != nil {
				log.Fatalf("Could not change urgency for id %d to %s", id, u.Arg)
			}
			project.ToFile(env.WorkDir)
		}
	} else {
		log.Fatal("Failed to get current working directory")
	}
}
