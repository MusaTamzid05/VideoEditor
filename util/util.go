package util

import (
	"os/exec"
	"strings"
)

func ExecuteCommand(commandStr string) (string, error) {
	data := strings.Split(commandStr, " ")
	command := data[0]
	args := data[1:]

	cmd := exec.Command(command, args...)
	stdout, err := cmd.Output()

	return string(stdout), err

}
