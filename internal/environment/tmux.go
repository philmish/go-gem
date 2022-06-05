package environment

import (
	"fmt"
	"os"
)

const SHEBANG = "#!/usr/bin/bash"

type Window struct {
	Name string `json:"name"`
	Cmd  string `json:"cmd"`
	Num  int    `json:"wnum"`
}

func (w *Window) createCmd(sName string) string {
	var cmd string
	if w.Num == 1 {
		cmd = fmt.Sprintf("tmux rename-window -t %s:%d -n '%s'", sName, w.Num, w.Name)
	} else {
		cmd = fmt.Sprintf("tmux new-window -t %s:%d -n '%s'", sName, w.Num, w.Name)
	}
	if w.Cmd != "" {
		cmd += fmt.Sprintf("\ntmux send-keys -t %s:%d '%s'", sName, w.Num, w.Cmd)
	}
	return cmd
}

func NewWindow(name, cmd string, num int) *Window {
	return &Window{Name: name, Cmd: cmd, Num: num}
}

type Session struct {
	Name    string   `json:"name"`
	Windows []Window `json:"windows"`
}

func (s *Session) createCmd() string {
	cmd := fmt.Sprintf("tmux new -d -s %s", s.Name)
	return cmd
}

func NewSession(name string, windows []Window) *Session {
	return &Session{Name: name, Windows: windows}
}

func (s *Session) generate() []string {
	var session []string = []string{s.createCmd()}
	for _, window := range s.Windows {
		session = append(session, window.createCmd(s.Name))
	}
	return session
}

type Config struct {
	Name     string    `json:"name"`
	Sessions []Session `json:"sessions"`
}

func (c *Config) Generate() []string {
	var result []string = []string{SHEBANG}
	for _, session := range c.Sessions {
		for _, window := range session.generate() {
			result = append(result, window)
		}
	}
	return result
}

func NewConfig(name string, sessions []Session) *Config {
	return &Config{Name: name, Sessions: sessions}
}

func (c *Config) Write(path string) error {
	var err error
	conf := c.Generate()
	fPath := fmt.Sprintf("%s/%s", path, c.Name)
	f, err := os.Create(fPath)
	if err != nil {
		f.Close()
		return err
	}

	for _, v := range conf {
		_, err = fmt.Fprintln(f, v)
		if err != nil {
			f.Close()
			return err
		}
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return err
}
