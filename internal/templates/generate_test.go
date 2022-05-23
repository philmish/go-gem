package templates_test

import (
	"github.com/philmish/go-gem/internal/environment"
	"github.com/philmish/go-gem/internal/templates"
	"testing"
)

func checkErr(err error, t *testing.T) {
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestDefaults(t *testing.T) {
	names := []string{"go", "node", "vue"}
	for _, name := range names {
		_, err := templates.CreateTemplate(name, "/home/user", true)
		checkErr(err, t)
	}
}
