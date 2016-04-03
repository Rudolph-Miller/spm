package command

import (
	"fmt"
	"github.com/Rudolph-Miller/spm/util"
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
	exists, _ := spm_util.Exists(target)
	if exists {
		fmt.Printf("%s is already installed.\n", name)
	} else {
		out, err := exec.Command("git", "clone", repo, target).CombinedOutput()
		fmt.Println(string(out))
		if err != nil {
			os.Exit(1)
		}
	}
}
