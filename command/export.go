package command

import (
	"fmt"
	"github.com/codegangsta/cli"
)

func CmdExport(c *cli.Context, directory string) {
	fmt.Println(c.String("output"))
}
