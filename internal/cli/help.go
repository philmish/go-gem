package cli

import "fmt"

func helpCmdParser(command string) {
        switch command {
                case "init":
                        helpInit()
                case "ls":
                        helpLs()
                case "add":
                        helpAdd()
                default:
                        helpCmd()
                }
}

func helpCmd() {
    fmt.Println("-------------------- HELP --------------------")
	fmt.Println("go-gem is a general environment manager for Linux.")
    fmt.Println("It uses a json config file to manage custom shell commands and todos for a Project.")
    fmt.Println("Commands can be saved to a alias file which you just need to source to use the save commands in your terminal.")
    fmt.Println("")

    fmt.Println("Usage:")
    fmt.Println("\tgogem -c help [command]")
    fmt.Println("\tgogem -c init [-n <template name>] [--alias]")
    fmt.Println("\tgogem -c ls")
    fmt.Println("\tgogem -c add -n <cmd alias> -a <shell cmd> [<args>...]")
    fmt.Println("\tgogem -c do -n <cmd alias>")
    fmt.Println("\tgogem -c rm -n <cmd alias>")
    fmt.Println("\tgogem -c lstodo")
    fmt.Println("\tgogem -c lsdone")
    fmt.Println("\tgogem -c addtodo -n <todo content> -a <todo urgency>")
    fmt.Println("\tgogem -c deltodo -n <todo id>")
    fmt.Println("\tgogem -c churg -n <todo id> -a <new urgency>")
    fmt.Println("")

    fmt.Println("Options:")
    fmt.Println("\t--alias\t If set go-gem creates a .gem_aliases file with aliases for all commands.")

}

func helpInit() {
        fmt.Println("-------------------- INIT --------------------")
        fmt.Println("The init command creates a new project config file in the current directory.")
        fmt.Println("go-gem provides templates for python, go and node projects.")
        fmt.Println("When no name or an unknown template name is provided go-gem creates a default project.")
        fmt.Println("")
        fmt.Println("Usage:")
        fmt.Println("\tgogem -c init [--alias]")
        fmt.Println("\tgogem -c init [-n go] [--alias]")
        fmt.Println("\tgogem -c init [-n py] [--alias]")
        fmt.Println("\tgogem -c init [-n node] [--alias]")
        fmt.Println("")

        fmt.Println("Options:")
        fmt.Println("\t--alias\t If set go-gem creates a .gem_aliases file with aliases for all commands.")
}

func helpLs() {
        fmt.Println("-------------------- LS --------------------")
        fmt.Println("The ls command lists all registered commands with their corresponding alias and shell command.")
        fmt.Println("")

        fmt.Println("Usage:")
        fmt.Println("\tgogem -c ls")
}

func helpAdd() {
        fmt.Println("-------------------- ADD --------------------")
        fmt.Println("The add command adds an alias and an command to your gem_conf.json.")
        fmt.Println("If you provided the --alias flag on project creation adding a new command will update your .gem_aliases file.")
        fmt.Println("If you want to add an arg for a command which contains whitespaces put the arg in double quotes.")
        fmt.Println("")

        fmt.Println("Usage:")
        fmt.Println("\tgogem -c add -n <alias> -a <shell command> <args>...")
}

func helpDo() {
        fmt.Println("-------------------- DO --------------------")
        fmt.Println("The do command invokes the shell command saved under the provided alias.")
        fmt.Println("If alias yout project it might be easier to source the .gem_aliases and use the command aliases directly.")
        fmt.Println("If you want you can provide additional arguments to the shell command.")
        fmt.Println("")

        fmt.Println("Usage:")
        fmt.Println("\tgogem -c do -n <alias> [<args>...]")
}

func helpRm() {
        fmt.Println("-------------------- RM --------------------")
        fmt.Println("The rm command removes an alias from the gem_conf.json.")
        fmt.Println("If you alias your project the rm command also removes the alias from the .gem_alias file.")
        fmt.Print("")

        fmt.Println("Usage:")
        fmt.Println("\tgogem -c rm -n <alias>")
}
