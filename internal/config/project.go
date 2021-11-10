package config

import (
	"encoding/json"
	"github.com/philmish/go-gem/internal/environment"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Project struct {
	Name  string                  `json:"name"`
	Todos []string                `json:"todos"`
	Env   environment.Environment `json:"env"`
}

func NewProject() *Project {
	var env = environment.NewEnv("")
	return &Project{
		"",
		[]string{},
		*env,
	}
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
