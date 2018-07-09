package modpack

import "path/filepath"

func (i *Instance) Base() string {
	return i.BaseDir
}

func (i *Instance) BinDir() string {
	return filepath.Join(i.Base(), "bin")
}

func (i *Instance) InstallDir() string {
	return filepath.Join(i.Base(), i.Version)
}
