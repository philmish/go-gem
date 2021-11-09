package cli

import (
        "flag"
)

type UserInput struct {
        Cmd string
        Name string
        Arg string
        AddArgs []string
}

func GetInput() *UserInput{
        var (
                cmd string
                name string
                arg string
        )

        flag.StringVar(&cmd, "c", "help", "gem command you want to run.")
        flag.StringVar(&name, "n", "", "gem module to use for parsing the args.")
        flag.StringVar(&arg, "a", "", "argument passed to the gem command.")
        flag.Parse()

        var data = &UserInput{cmd, name, arg, flag.Args()}
        return data
}



