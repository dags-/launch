package home

import (
	"github.com/dags-/launch/util/config"
	"github.com/dags-/launch/util/log"
	"github.com/zserge/webview"
)

type Home struct {
	wv       webview.WebView
	Modpacks []*Modpack `json:"modpacks"`
}

type Modpack struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Version  string `json:"version"`
	Selected bool   `json:"selected"`
}

var (
	home = &Home{
		Modpacks: []*Modpack{
			{Name: "AC Modpack", Version: "1.2.1", ID: "acmodpack", Selected: true},
			{Name: "CR Modpack", Version: "2.2.4", ID: "crmodpack", Selected: false},
		},
	}
)

func GetHome() *Home {
	return home
}

func (h *Home) Attach(w webview.WebView) {
	h.wv = w
}

func (h *Home) SetSize(width, height int) {
	config.SetSize(width, height)
}

func (h *Home) Action(name string) {
	log.Info("action: %s", name)
}

func (h *Home) Launch() {
	log.Info("launching....")
}

func (h *Home) SelectPack(id string) {
	log.Debug("selecting pack %s", id)
	for _, pack := range h.Modpacks {
		if pack.ID == id {
			log.Info("selected pack %s", id)
			pack.Selected = true
		} else {
			pack.Selected = false
		}
	}
}
