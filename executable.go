package common

import "os/exec"

func HasExecutable(binPath string) (string, error) {
	exec.LookPath(binPath)
}
