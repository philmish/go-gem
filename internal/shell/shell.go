package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/philmish/go-gem/internal/config"
	"github.com/philmish/go-gem/internal/parser"
)

type GemShell struct {
    Project config.Project
}

func parse_input(input string) *parser.UserInput {
    parts := strings.Split(input, " ")
    flags := map[string]string{
        "-n": "",
    }
    index := 0
    expected := ""
    var parsed *parser.UserInput
    for _, word := range parts {
        switch word {
        case "-n":
            expected = word
            continue
        default:
            break
        }

        if _, ok := flags[expected]; ok {
            flags[expected] = word
            continue
        }


        switch index {
        case 0:
            parsed.Cmd = word
            index++
            continue
        case 1:
            parsed.Name = word
            index++
            continue
        case 2:
            parsed.Arg = word
            index++
            continue
        default:
            parsed.AddArgs = append(parsed.AddArgs, word)
            index++
            continue
        }
        
    }
    return parsed
}


func (sh GemShell) RunShell() {
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("> " )
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
        input = strings.TrimSuffix(input, "\n")

        cmd := parse_input(input)
        parser.ParseEnvCommand(&sh.Project, cmd)
        // Placeholder shell piping input to stdin
        // TODO implemenet shell based command running
        /*
        cmd := exec.Command(input)
        cmd.Stderr = os.Stderr
        cmd.Stdout = os.Stdout
        err = cmd.Run()
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
        */
    }
}
