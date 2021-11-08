package cli

import (
        "flag"
)

type UserInput struct {
        Cmd string
        Module string
        Arg string
        AddArgs []string
}

func GetInput() *UserInput{
        var (
                cmd string
                module string
                arg string
        )

        flag.StringVar(&cmd, "cmd", "help", "gem command you want to run.")
        flag.StringVar(&module, "mod", "", "gem module to use for parsing the args.")
        flag.StringVar(&arg, "arg", "", "arg for the module")
        flag.Parse()

        var data = &UserInput{cmd, module, arg, flag.Args()}
        return data
}



