package main

import (
	"fmt"
	"github.com/Rudolph-Miller/spm/command"
	"github.com/codegangsta/cli"
	"os"
	"os/user"
	"strings"
)

var directory string

var default_directories = []string{
	"~/Library/Containers/com.bohemiancoding.sketch3/Data/Library/Application Support/sketch/Plugins/",
	"~/Library/Application Support/com.bohemiancoding.sketch3/Plugins/",
}

var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name:        "dir, d",
		EnvVar:      "SKETCH_PLUGIN_DIR",
		Usage:       "directory for Sketch plugins",
		Destination: &directory,
	},
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func Directory() string {
	if len(directory) > 0 {
		return directory
	} else {
		for _, dir := range default_directories {
			usr, _ := user.Current()
			path := strings.Replace(dir, "~", usr.HomeDir, 1)
			exists, _ := exists(path)
			if exists {
				directory = path
				return directory
			}
		}
	}
	fmt.Fprintf(os.Stderr, "Could not find Sketch plugins directory.\nPlease specify --dir option or SKETCH_PLUGIN_DIR env var.")
	os.Exit(1)
	return ""
}

var Commands = []cli.Command{
	{
		Name:    "install",
		Aliases: []string{"i"},
		Usage:   "Install plugin",
		Action: func(c *cli.Context) {
			command.CmdInstall(c, Directory())
		},
		Flags: []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.\n", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
