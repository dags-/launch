package modpack

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dags-/launch/util"
	"github.com/dags-/launch/util/log"
)

func Install(pack *Dependency, baseDir string, experimental bool) (*Instance, error) {
	return InstallVersion(pack, baseDir, "", experimental)
}

func InstallVersion(pack *Dependency, baseDir, version string, experimental bool) (*Instance, error) {
	i := &Instance{
		Dependency: &Dependency{
			Version: version,
			Repo: Repo{
				Owner: pack.Owner,
				Name:  pack.Name,
			},
		},
		Experimental: experimental,
		BaseDir:      baseDir,
		Parents:      make([]*Dependency, 0),
		Mods:         make(map[string]Mod),
	}

	defer tidy(i.BinDir(), ".zip")

	if version == "" {
		return i, i.installLatest(pack)
	}

	return i, i.installVersion(pack, version)
}

func (i *Instance) Update() (*Instance, error) {
	rel, err := i.GetLatest(i.Experimental)
	if err != nil {
		return nil, err
	}

	if rel.Tag == i.Version {
		log.Info("no update necessary")
		return i, nil
	}

	return InstallVersion(i.Dependency, i.BaseDir, rel.Tag, i.Experimental)
}

func tidy(dir, ext string) error {
	files, e := ioutil.ReadDir(dir)
	if e != nil {
		return e
	}
	for _, f := range files {
		if filepath.Ext(f.Name()) == ext {
			os.Remove(f.Name())
		}
	}
	return nil
}

func (i Instance) installLatest(p *Dependency) error {
	rel, err := p.GetLatest(i.Experimental)
	if err != nil {
		return err
	}
	i.Dependency = &Dependency{
		Version: rel.Tag,
		Repo: Repo{
			Owner: p.Owner,
			Name:  p.Name,
		},
	}
	return i.install(rel.URL)
}

func (i *Instance) installVersion(p *Dependency, version string) error {
	rel, err := p.GetRelease(version)
	if err != nil {
		return err
	}
	return i.install(rel.URL)
}

func (i *Instance) install(url string) error {
	name := fmt.Sprintf("%v.zip", time.Now().Unix())
	z, err := util.Download(url, filepath.Join(i.BaseDir, "bin"), name)
	if err != nil {
		return err
	}

	var pack Meta
	err = util.Extract(z, filepath.Join(i.BaseDir, i.Version), pathFunc(&pack))
	if err != nil {
		return err
	}

	if pack.Parent != nil && pack.Parent.ID() != i.ID() {
		i.Parents = append(i.Parents, pack.Parent)
		i.installVersion(pack.Parent, pack.Parent.Version)
	}

	if pack.Mods != nil {
		for k, v := range *pack.Mods {
			i.Mods[k] = v
		}
	}

	return nil
}

func pathFunc(mp *Meta) func(file *zip.File) string {
	return func(zf *zip.File) string {
		parts := strings.Split(zf.Name, "/")
		if len(parts) > 1 {
			p := filepath.Join(parts[1:]...)
			if p == "modpack.json" {
				in, e := zf.Open()
				if e == nil {
					defer in.Close()
					e = json.NewDecoder(in).Decode(mp)
				}
				return ""
			}
			return p
		}
		return ""
	}
}
