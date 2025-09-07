package hledger

import (
	"bytes"
	"io"
	"os/exec"
	"strings"
)

func (h Hledger) Custom(command string) (io.Reader, error) {
	args := strings.Split(command, " ")

	if h.logger != nil {
		h.logger.Printf("running command: %s %s", h.hledgerBinary, strings.Join(args, " "))
	}

	cmd := exec.Command(h.hledgerBinary, args...)
	result, err := cmd.CombinedOutput()
	if err != nil {
		return bytes.NewBuffer(result), &Error{cmdArgs: args, err: err}
	}

	return bytes.NewBuffer(result), nil
}
