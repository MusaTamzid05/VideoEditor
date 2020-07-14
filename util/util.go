package util

import (
	"fmt"
	"os"
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

func GenereateDstFiles(outputFile string, paths []string) error {

	// delete the output file if exists

	if Exists(outputFile) {
		os.Remove(outputFile)
	}

	fp, err := os.Create(outputFile)

	if err != nil {
		return err
	}

	defer fp.Close()

	newPaths := []string{}

	for _, path := range paths {
		newPaths = append(newPaths, fmt.Sprintf("file %s", path))
	}

	_, err = fp.WriteString(strings.Join(newPaths, "\n"))

	if err != nil {
		return err
	}

	return nil
}

func Exists(path string) bool {

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
