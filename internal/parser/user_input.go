package parser

import (
	"flag"
    "strings"
)

type UserInput struct {
	Cmd     string
	Name    string
	Arg     string
	Alias   bool
	AddArgs []string
}

func GetInput() *UserInput {
	var (
		cmd   string
		name  string
		arg   string
		alias bool
	)

	flag.StringVar(&cmd, "c", "help", "gem command you want to run.")
	flag.StringVar(&name, "n", "", "gem module to use for parsing the args.")
	flag.StringVar(&arg, "a", "", "argument passed to the gem command.")
	flag.BoolVar(&alias, "alias", false, "create an alias file.")
	flag.Parse()

	var data = &UserInput{cmd, name, arg, alias, flag.Args()}
	return data
}

func InputFromString(input string) *UserInput {
    parts := strings.Split(input, " ")
    flags := map[string]string{
        "-n": "",
    }
    index := 0
    expected := ""
    var parsed *UserInput
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
