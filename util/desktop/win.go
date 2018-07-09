package desktop

import "os/exec"

type windows struct {
}

func (w *windows) Open(path string) error {
	return exec.Command("cmd", "/C", "start", path).Run()
}
