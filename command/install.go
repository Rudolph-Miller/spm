package command

import (
	"fmt"
	"github.com/codegangsta/cli"
)

func CmdInstall(c *cli.Context, directory string) {
	fmt.Printf("%s", directory)
}
