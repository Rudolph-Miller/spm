package spm_struct

import (
	"github.com/Rudolph-Miller/spm/util"
	"path/filepath"
)

type Source struct {
	Name string
	Path string
}

type SourceType int

const (
	Git = iota
	Other
)

func (s *Source) Type() SourceType {
	gitExists, _ := spm_util.Exists(filepath.Join(s.Path, ".git"))

	switch {
	case gitExists:
		return Git
	}
	return Other
}
