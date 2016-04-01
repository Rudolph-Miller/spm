package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
	"path/filepath"
)

func CmdInstall(c *cli.Context, directory string) {
	args := c.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Please specify target.\n")
		os.Exit(1)
	}
	repo := args.First()
	_, name := filepath.Split(repo)
	target := filepath.Join(directory, name)
	out, err := exec.Command("git", "clone", repo, target).CombinedOutput()
	fmt.Println(string(out))
	if err != nil {
		os.Exit(1)
	}
}
