package templates

import (
        "github.com/philmish/go-gem/internal/environment"
)

func DefaultNodeEnv(dir string) *environment.Environment {
        var newEnv = environment.NewEnv(dir)

        newEnv.Add("v", "nvm", []string{"use"})
        newEnv.Add("vinst", "nvm", []string{"install"})
        newEnv.Add("get", "npm", []string{"install"})

        return newEnv
}
