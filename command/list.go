package command

import (
	"fmt"
	"github.com/Rudolph-Miller/spm/struct"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CmdList(c *cli.Context, directory string) {
	files, err := ioutil.ReadDir(directory)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		name := file.Name()
		source := spm_struct.Source{
			Name: name,
			Path: filepath.Join(directory, name),
		}
		switch source.Type() {
		case spm_struct.Git:
			fmt.Printf("[Git] ")
		case spm_struct.Other:
			fmt.Printf("[Other] ")
		}
		fmt.Println(source.Name)
	}
}
