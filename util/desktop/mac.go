package desktop

import "os/exec"

type mac struct {
}

func (m *mac) Open(path string) error {
	return exec.Command("open", path).Run()
}
