package command

import (
	"fmt"
	"github.com/Rudolph-Miller/spm/util"
	"github.com/codegangsta/cli"
	"os"
	"path/filepath"
)

func CmdUninstall(c *cli.Context, directory string) {
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
		err := os.RemoveAll(target)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("%s is not found.\n", name)
	}
}
