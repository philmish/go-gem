package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RunShell() {
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("> " )
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
        input = strings.TrimSuffix(input, "\n")
        // Placeholder shell piping input to stdin
        // TODO implemenet shell based command running
        cmd := exec.Command(input)
        cmd.Stderr = os.Stderr
        cmd.Stdout = os.Stdout
        err = cmd.Run()
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
    }
}
