// +build freebsd linux netbsd openbsd solaris dragonfly

package open

import "os/exec"

var (
	openCmds = []string{"xdg-open", "cygstart", "x-www-browser", "firefox", "opera", "mozilla", "netscape"}
)

func open(args ...string) error {
	var cmd string
	for _, cmd = range openCmds {
		path, err := exec.LookPath(cmd)
		if err == nil {
			cmd = path
			break
		}
	}

	if err := exec.Command(cmd, args...).Start(); err != nil {
		return err
	}
	return nil
}
