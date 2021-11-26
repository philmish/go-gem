package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/philmish/go-gem/internal/environment"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
)

type Todo struct {
	Content string `json:"content"`
	Done    bool   `json:"done"`
	Urgency int8   `json:"urgency"`
	Id      int    `json:"id"`
}

func newTodo(content string, urgency int8, id int) *Todo {
	return &Todo{content, false, urgency, id}
}

func (t *Todo) formatTodo() string {
	return fmt.Sprintf("(%d)%s\nID: %d\n\n", t.Urgency, t.Content, t.Id)
}

type Project struct {
	Name  string                  `json:"name"`
	Todos []Todo                  `json:"todos"`
	Env   environment.Environment `json:"env"`
}

func NewProject() *Project {
	var env = environment.NewEnv("")
	return &Project{
		"",
		[]Todo{},
		*env,
	}
}

func (p *Project) checkTodoId(id int) (int, error) {
	for i, item := range p.Todos {
		if item.Id == id {
			return i, nil
		}
	}
	return 0, errors.New("No todo found for id")
}

func (p *Project) checkForTodoContent(content string) *Todo {
	for _, i := range p.Todos {
		if i.Content == content {
			return &i
		}
	}
	return nil
}

func (p *Project) AddTodo(content string, urgency int8) error {
	exists := p.checkForTodoContent(content)
	if exists != nil {
		return errors.New("A todo with the same content already exists.")
	}
	nId := len(p.Todos) + 1
	nTodo := newTodo(content, urgency, nId)
	p.Todos = append(p.Todos, *nTodo)
	return nil
}

func (p *Project) DelTodo(id int) error {
	index, err := p.checkTodoId(id)
	if err != nil {
		return err
	}
	p.Todos[index].Done = true
	return nil
}

func (p *Project) ChangeUrgency(id int, urgency int8) error {
	index, err := p.checkTodoId(id)
	if err != nil {
		return err
	}
	p.Todos[index].Urgency = urgency
	return nil
}

func (p *Project) ListTodos(done bool) string {
	sort.Slice(p.Todos, func(i, j int) bool { return p.Todos[i].Urgency > p.Todos[j].Urgency })
	var result string
	for _, i := range p.Todos {
		if i.Done == done {
			result += i.formatTodo()
		}
	}
	return result
}

func readFile(fpath string, project *Project) error {
	conf_file := filepath.Join(fpath, "gem_config.json")
	if _, err := os.Stat(conf_file); err == nil {
		jf, err := os.Open(conf_file)
		if err != nil {
			return err
		}
		defer jf.Close()
		byteValue, _ := ioutil.ReadAll(jf)
		err = json.Unmarshal(byteValue, project)
		if err != nil {
			return err
		}
		return nil
	} else {
		return err
	}
}

func FromFile(fpath string) *Project {
	data := NewProject()
	err := readFile(fpath, data)

	if err != nil {
		log.Fatalf("Could not read %s\nError: %v", fpath, err)
	}

	return data
}

func (p *Project) ToFile(fpath string) error {
	fp := filepath.Join(fpath, "gem_config.json")
	file, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		return err
	}

	if p.Env.Alias {
		err = p.Env.Aliases()
		if err != nil {
			return err
		}
	}
	err = ioutil.WriteFile(fp, file, 0644)
	return err
}
