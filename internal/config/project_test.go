package config_test

import (
	"fmt"
	"github.com/philmish/go-gem/internal/config"
	"os"
	"testing"
)

func checkErr(err error, t *testing.T) {
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestProject(t *testing.T) {
	p, err := os.Getwd()

	if err != nil {
		t.Error("Failed to get working directory.")

	}

	project := config.NewProject()
	project.Name = "TestProject"
	project.Env.WorkDir = p
	project.AddTodo("Project Test Todo", 5)
	project.DelTodo(1)

	if !project.Todos[0].Done {
		t.Error("Tried to delete todo but failed.")
	}

	var env = project.Env
	env.Add("hello", "echo", []string{"Hello World"})
	env.List()
	env.Do("hello", []string{"with", "args"})

	err = project.ToFile(p)

	if err != nil {
		t.Error("Failed to write Project to file.")
	}

	loadedConf := config.FromFile(p)

	fmt.Println(loadedConf.Name)
}

func TestTodo(t *testing.T) {
	project := config.NewProject()
	project.Name = "TestProject"
	err := project.AddTodo("Test the todo function", 10)
	checkErr(err, t)
	todos := project.ListTodos()
	fmt.Printf(todos)
	err = project.ChangeUrgency(1, 12)
	checkErr(err, t)
	err = project.DelTodo(1)
	checkErr(err, t)
}
