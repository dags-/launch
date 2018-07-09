package modpack

import (
	"fmt"
	"path/filepath"

	"github.com/dags-/launch/util"
)

func (i *Instance) PreLaunch() {
	mods := i.compileDependencies()
	for _, mod := range mods {
		path := filepath.Join(i.InstallDir(), mod.Path)
		exists := util.Exists(path)
		if mod.Enabled && !exists {
			from := filepath.Join(i.BinDir(), mod.Path)
			to := filepath.Join(i.InstallDir(), mod.Path)
			util.Move(from, to)
		}
		if exists && !mod.Enabled {
			from := filepath.Join(i.InstallDir(), mod.Path)
			to := filepath.Join(i.BinDir(), mod.Path)
			util.Move(from, to)
		}
		fmt.Println(mod.Name, mod.Enabled)
	}
}

func (i *Instance) compileDependencies() map[string]Mod {
	compiled := make(map[string]Mod)

	for name, mod := range i.Mods {
		if mod.Dependencies == nil {
			compiled[name] = mod
			continue
		}
		if !mod.Enabled {
			compiled[name] = mod
			continue
		}

		for dependency, version := range *mod.Dependencies {
			if m, ok := compiled[dependency]; ok {
				if m.Enabled && m.Version == version {
					continue
				}
			} else if m, ok := i.Mods[dependency]; ok {
				if m.Enabled && m.Version == version {
					continue
				}
			}
			fmt.Println("disabling:", mod.Name, "due to missing/disabled dependency:", dependency)
			compiled[name] = Mod{
				Enabled: false,
				Name:    mod.Name,
				Path:    mod.Path,
			}
			break
		}

		if _, ok := compiled[name]; !ok {
			compiled[name] = mod
		}
	}

	return compiled
}
