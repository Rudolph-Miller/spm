package command

import (
	"fmt"
	"github.com/codegangsta/cli"
)

func CmdList(c *cli.Context, directory string) {
	fmt.Println(directory)
}
