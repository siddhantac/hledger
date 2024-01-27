package hledger

import (
	"bytes"
	"io"
	"os/exec"
)

func (h Hledger) execCmd(hledgerCmd string, options Options) (io.Reader, *Error) {
	args := make([]string, 0)
	args = append(args, hledgerCmd)

	if h.journalFilename != "" {
		args = append(args, "-f", h.journalFilename)
	}

	args = append(args, options.Build()...)

	// log.Printf("running command: %s %s", h.hledgerBinary, strings.Join(args, " "))

	cmd := exec.Command(h.hledgerBinary, args...)
	result, err := cmd.CombinedOutput()
	if err != nil {
		return bytes.NewBuffer(result), &Error{cmdArgs: args, err: err}
	}

	return bytes.NewBuffer(result), nil
}
