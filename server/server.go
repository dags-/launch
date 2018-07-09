package server

import (
	"net"
	"net/http"
	"time"

	"github.com/GeertJohan/go.rice"
	"github.com/dags-/launch/util/log"
	"github.com/dags-/launch/view"
	"github.com/dags-/launch/view/home"
	"github.com/gorilla/mux"
)

var (
	bind chan view.View
)

func Serve(bx *rice.Box) (net.Listener, chan view.View, error) {
	l, e := port()
	if e != nil {
		return nil, nil, e
	}

	bind = make(chan view.View, 1)
	r := mux.NewRouter()
	r.HandleFunc("/view/{view}/", ready)
	r.PathPrefix("/").Handler(http.FileServer(bx.HTTPBox()))

	go func(r *mux.Router, l net.Listener) {
		http.ListenAndServe(l.Addr().String(), r)
	}(r, l)

	log.Info("serving on %s", l.Addr())

	return l, bind, nil
}

func ready(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	time.Sleep(250 * time.Millisecond)
	log.Debug("binding view to client")
	switch vars["view"] {
	case "home":
		bind <- home.GetHome()
		return
	}
}

func port() (net.Listener, error) {
	lsn, e := net.Listen("tcp", "127.0.0.1:0")
	if e != nil {
		return nil, e
	}
	lsn.Close()
	return lsn, nil
}
