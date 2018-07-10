package main

import (
	"fmt"

	"github.com/GeertJohan/go.rice"
	"github.com/dags-/launch/server"
	"github.com/dags-/launch/util/config"
	"github.com/dags-/launch/util/log"
	"github.com/dags-/launch/view"
	"github.com/zserge/webview"
)

func main() {
	box := rice.MustFindBox("assets")

	log.Setup(config.Debug())
	l, c, e := server.Serve(box)
	if e != nil {
		log.Err("serve error %s", e)
		panic(e)
	}

	w, h := config.WindowSize()
	wv := webview.New(webview.Settings{
		URL:       fmt.Sprintf("http://%s/home", l.Addr()),
		Title:     "Launcher",
		Width:     w,
		Height:    h,
		Resizable: true,
	})

	go handleBind(wv, c)

	wv.Run()
	config.Save()
}

func handleBind(w webview.WebView, c chan view.View) {
	for {
		v := <-c
		v.Attach(w)

		w.Dispatch(func() {
			w.Bind("view", v)
			w.Eval(`onbind()`)
		})
	}
}
