# Sketch Plugin Manager

It can manage your plugins for Sketch.

## Description

## Usage

```
NAME:
   spm - 

USAGE:
   commands [global options] command [command options] [arguments...]
   
VERSION:
   0.1.0
   
AUTHOR(S):
   Rudolph-Miller 
   
COMMANDS:
   info   Information for Sketch plugins.
   directory, d Directory for Sketch plugins.
   list, l  List plugins.
   install, i Install plugin.
   uninstall, u Uninstall plugin.
   export, e  Export plugins into spmfile
   help, h  Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --dir, -d    directory for Sketch plugins [$SKETCH_PLUGIN_DIR]
   --help, -h   show help
   --version, -v  print the version
```

## Install

To install, use `go get`:

```bash
$ go get -d github.com/Rudolph-Miller/spm
```

## Commands

### info

### directory, d

```
$ spm d
/Users/sketch-user/Library/Application\ Support/com.bohemiancoding.sketch3/Plugins/

$ spm d | xargs mv ~/Downloads/AwesomePlugin.sketchplugin
```

### list, l


### install, i

### uninstall, u

### export, e

### help, h

## Contribution

1. Fork ([https://github.com/Rudolph-Miller/spm/fork](https://github.com/Rudolph-Miller/spm/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[Rudolph-Miller](https://github.com/Rudolph-Miller)
