package main

import (
	"fmt"
	"github.com/Rudolph-Miller/spm/command"
	"github.com/Rudolph-Miller/spm/util"
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

func Directory() string {
	if len(directory) > 0 {
		return directory
	} else {
		for _, dir := range default_directories {
			usr, _ := user.Current()
			path := strings.Replace(dir, "~", usr.HomeDir, 1)
			exists, _ := spm_util.Exists(path)
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
		Name:  "info",
		Usage: "Information for Sketch plugins.",
		Action: func(c *cli.Context) {
			directory = strings.Replace(Directory(), " ", "\\ ", -1)
			fmt.Printf("Directory: %s\n", directory)
		},
		Flags: []cli.Flag{},
	},
	{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "List plugins.",
		Action: func(c *cli.Context) {
			command.CmdList(c, Directory())
		},
		Flags: []cli.Flag{},
	},
	{
		Name:    "install",
		Aliases: []string{"i"},
		Usage:   "Install plugin.",
		Action: func(c *cli.Context) {
			command.CmdInstall(c, Directory())
		},
		Flags: []cli.Flag{},
	},
	{
		Name:    "uninstall",
		Aliases: []string{"u"},
		Usage:   "Uninstall plugin.",
		Action: func(c *cli.Context) {
			command.CmdUninstall(c, Directory())
		},
		Flags: []cli.Flag{},
	},
	{
		Name:    "export",
		Aliases: []string{"e"},
		Usage:   "Export plugins into spmfile",
		Action: func(c *cli.Context) {
			command.CmdExport(c, Directory())
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "output, o",
				Value:  "Spmfile",
				EnvVar: "SPM_OUTPUT_FILE",
				Usage:  "output file to export plugins.",
			},
		},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.\n", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
