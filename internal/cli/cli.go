package cli

import (
        "flag"
)

type UserInput struct {
        Cmd string
        Name string
        Arg string
        Alias bool
        AddArgs []string
}

func GetInput() *UserInput{
        var (
                cmd string
                name string
                arg string
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



