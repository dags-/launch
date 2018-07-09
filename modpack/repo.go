package modpack

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
)

var (
	TagNotFound      = errors.New("tag not found")
	ReleasesNotFound = errors.New("releases not found")
)

type Repo struct {
	Owner    string `json:"owner"`
	Name     string `json:"name"`
	lock     sync.RWMutex
	releases []*Release
}

type Release struct {
	Tag        string `json:"tag_name"`
	Changelog  string `json:"body"`
	PreRelease bool   `json:"prerelease"`
	URL        string `json:"zipball_url"`
}

func (r *Repo) Fetch() (error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", r.Owner, r.Name)
	rq, e := http.Get(url)
	if e != nil {
		return e
	}
	defer rq.Body.Close()

	var releases []*Release
	e = json.NewDecoder(rq.Body).Decode(&releases)
	if e != nil {
		return e
	}
	r.releases = releases
	return nil
}

func (r *Repo) GetLatest(experimental bool) (*Release, error) {
	if r.releases == nil {
		r.Fetch()
	}
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.releases == nil {
		return nil, ReleasesNotFound
	}
	for _, rl := range r.releases {
		if !rl.PreRelease || experimental {
			return rl, nil
		}
	}
	return nil, ReleasesNotFound
}

func (r *Repo) GetRelease(tag string) (*Release, error) {
	if r.releases == nil {
		r.Fetch()
	}
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.releases == nil {
		return nil, ReleasesNotFound
	}
	for _, rl := range r.releases {
		if rl.Tag == tag {
			return rl, nil
		}
	}
	return nil, TagNotFound
}
