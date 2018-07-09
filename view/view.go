package view

import "github.com/zserge/webview"

type View interface {
	Attach(wv webview.WebView)

	Action(name string)

	SetSize(width, height int)
}
