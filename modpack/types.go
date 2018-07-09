package modpack

import (
	"path/filepath"
)

type Manifest []Dependency

type Dependency struct {
	Repo
	Version string `json:"version,omitempty"`
}

type Meta struct {
	Parent *Dependency     `json:"parent,omitempty"`
	Mods   *map[string]Mod `json:"mods,omitempty"`
}

type Instance struct {
	*Dependency
	BaseDir      string         `json:"base_dir"`
	Experimental bool           `json:"experimental"`
	Parents      []*Dependency  `json:"parents"`
	Mods         map[string]Mod `json:"mods"`
}

type Mod struct {
	Name         string             `json:"name"`
	Version      string             `json:"version"`
	Path         string             `json:"path"`
	Enabled      bool               `json:"enabled"`
	Dependencies *map[string]string `json:"dependencies,omitempty"`
}

func (p *Dependency) ID() string {
	return filepath.Join(p.Owner, p.Name)
}
