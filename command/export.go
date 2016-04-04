package command

import (
	"bytes"
	"fmt"
	"github.com/Rudolph-Miller/spm/struct"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CmdExport(c *cli.Context, directory string) {
	files, err := ioutil.ReadDir(directory)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var buffer bytes.Buffer

	for _, file := range files {
		name := file.Name()
		source := spm_struct.Source{
			Name: name,
			Path: filepath.Join(directory, name),
		}
		switch source.Type() {
		case spm_struct.Git:
			buffer.Write([]byte("git"))
		case spm_struct.Other:
			buffer.Write([]byte("other"))
		}
		buffer.Write([]byte{' '})
		buffer.Write([]byte(source.Name))
		buffer.Write([]byte{'\n'})
	}
	fmt.Println(buffer.String())
}
