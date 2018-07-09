package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"sync"

	"github.com/dags-/launch/util"
	"github.com/dags-/launch/util/log"
)

var (
	l    sync.RWMutex
	cfg  = load()
	path = filepath.Join(workDir(), "config.json")
)

type Config struct {
	Debug        bool `json:"debug"`
	WindowWidth  int  `json:"window_width"`
	WindowHeight int  `json:"window_height"`
}

func SetSize(width, height int) {
	l.Lock()
	defer l.Unlock()
	cfg.WindowWidth = width
	cfg.WindowHeight = height
}

func Debug() bool {
	l.RLock()
	defer l.RUnlock()
	return cfg.Debug
}

func WindowSize() (int, int) {
	l.RLock()
	defer l.RUnlock()
	return cfg.WindowWidth, cfg.WindowHeight
}

func Save() {
	l.Lock()
	defer l.Unlock()
	write(*cfg)
}

func load() *Config {
	var c Config

	d, e := ioutil.ReadFile(path)
	if e == nil {
		e = json.Unmarshal(d, &c)
		if e == nil {
			return &c
		}
	}

	log.Info("config.load error %s", e)

	c = Config{
		Debug:        true,
		WindowWidth:  800,
		WindowHeight: 480,
	}

	write(c)

	return &c
}

func write(c Config) {
	log.Info("config.save writing %s", path)
	if !util.Exists(path) {
		os.MkdirAll(filepath.Dir(path), os.ModePerm)
	}

	d, e := json.MarshalIndent(c, "", "  ")
	if e == nil {
		e = ioutil.WriteFile(path, d, os.ModePerm)
	}

	if e != nil {
		log.Info("config.save error %s", e)
	}
}

func workDir() string {
	usr, err := user.Current()
	if err != nil {
		p, e := os.Executable()
		if e != nil {
			return ""
		}
		return p
	}
	return filepath.Join(usr.HomeDir, "Documents", "launcher")
}
