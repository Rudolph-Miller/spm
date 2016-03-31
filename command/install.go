package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
)

func CmdInstall(c *cli.Context, directory string) {
	args := c.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Please specify target.\n")
		os.Exit(1)
	}
	target := args.First()
	out, err := exec.Command("git", "clone", target, directory).CombinedOutput()
	fmt.Println(string(out))
	if err != nil {
		os.Exit(1)
	}
}
