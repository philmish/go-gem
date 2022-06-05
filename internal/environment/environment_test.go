package environment_test

import (
	"github.com/philmish/go-gem/internal/environment"
	"os"
	"testing"
)

func checkErr(e error, t *testing.T) {
	if e != nil {
		t.Errorf("%v", e)
	}
}

func contains(val string, target []string) bool {
	for _, v := range target {
		if v == val {
			return true
		}
	}
	return false
}

func TestCommand(t *testing.T) {
	newCmd := environment.Command{"echo", []string{"Hello World"}}
	newCmd.PrintCmd()
	newCmd.Execute()
}

func TestEnvironment(t *testing.T) {
	p, err := os.Getwd()
	checkErr(err, t)
	var newEnv = environment.NewEnv(p)

	newEnv.Add("hello", "echo", []string{"Hello World"})
	newEnv.List()
	newEnv.Do("hello", []string{"with", "args"})
	err = newEnv.Aliases()
	checkErr(err, t)
	newEnv.Remove("hello")
	newEnv.List()
}

func TestTmux(t *testing.T) {
	// Test tmux
	windows := []environment.Window{
		{Name: "testWdw1", Cmd: "nvim", Num: 1},
		{Name: "testWdw2", Cmd: "", Num: 2},
		{Name: "s2testWdw1", Cmd: "whoami", Num: 1},
		{Name: "s2testWdw2", Cmd: "", Num: 1},
	}
	sessions := []environment.Session{
		{Name: "testSesh1", Windows: []environment.Window{windows[0], windows[1]}},
		{Name: "testSesh2", Windows: []environment.Window{windows[2], windows[3]}},
	}
	conf := environment.NewConfig("testConf", sessions)
	commands := conf.Generate()

	testVals := []string{"#!/usr/bin/bash", "tmux new -d -s testSesh1", "tmux rename-window -t testSesh1:1 -n 'testWdw1'\ntmux send-keys -t testSesh1:1 'nvim'"}
	for _, tVal := range testVals {
		exists := contains(tVal, commands)
		if !exists {
			t.Errorf("Commands missing: %s\n", tVal)
		}
	}
}
