package lib

import (
	"log"
	"os/exec"
	"strings"
)

func ExecCmd(cmd string) string {
	fields := strings.Fields(cmd)
	bytes, err := exec.Command(fields[0], fields[1:]...).Output()
	s := strings.TrimSpace(string(bytes))
	if err != nil {
		log.Fatal(err)
	}
	return s
}
