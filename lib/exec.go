package lib

import (
	"os/exec"
	"strings"
)

func ExecCmd(cmd string) (string, error) {
	fields := strings.Fields(cmd)
	bytes, err := exec.Command(fields[0], fields[1:]...).Output()
	if err != nil {
		return "", err
	}
	// else
	s := strings.TrimSpace(string(bytes))
	return s, nil
}
