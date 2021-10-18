// +build windows
package open

import "os/exec"

const (
	openCmd = "cmd /c start"
)

func getOpenCommand(args ...string) *exec.Cmd {
	return exec.Command(openCmd, args...)
}

func open(args ...string) error {
	cmd := getOpenCommand(args...)
	if err := cmd.Start(); err != nil {
		return err
	}
	return nil
}
