package desktop

import "runtime"

type Platform interface {
	Open(string) error
}

var platform Platform

func init() {
	switch runtime.GOOS {
	case "windows":
		platform = &windows{}
	case "darwin":

	case "linux":
	}
}

func Open(path string) error {
	return platform.Open(path)
}
