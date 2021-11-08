package config_test

import (
        "github.com/philmish/go-gem/internal/config"
        "testing"
        "os"
        "fmt"
)

func TestProject(t *testing.T){
        p, err := os.Getwd()

        if err != nil{
                t.Error("Failed to get working directory.")

        }

        project := config.NewProject()
        project.Name = "TestProject"
        project.Env.WorkDir = p
        project.Todos = append(project.Todos, "Test todo")
        
        var env = project.Env
        env.Add("Setup", "hello", "echo", []string{"Hello World"})
        env.List("Setup")
        env.Do("Setup", "hello")

        err = project.ToFile(p)

        if err != nil{
                t.Error("Failed to write Project to file.")
        }

        loadedConf := config.FromFile(p)

        fmt.Println(loadedConf.Name)
}
