package hledger

import (
	"bytes"
	"io"
	"log"
	"os/exec"
	"strings"
)

type Filter interface {
	Build() string
}

type NoFilter struct{}

func (NoFilter) Build() string { return "" }

func (h Hledger) execCmd(hledgerArgs []string, filters ...Filter) (io.Reader, error) {
	args := h.buildCmd(hledgerArgs, filters...)
	log.Printf("running command: %s", strings.Join(args, " "))

	cmd := exec.Command(h.HledgerBinary, args...)
	result, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("error: %v", err)
		return bytes.NewBuffer(result), err
	}

	return bytes.NewBuffer(result), nil
}

func (h Hledger) buildCmd(hledgerArgs []string, filters ...Filter) []string {
	args := make([]string, 0)

	args = append(args, hledgerArgs...)

	for _, f := range filters {
		if _, ok := f.(NoFilter); ok {
			continue
		}
		args = append(args, f.Build())
	}

	if h.JournalFilename != "" {
		args = append(args, "-f", h.JournalFilename)
	}

	args = append(args, "-O", "csv")

	return args
}

func (h Hledger) execWithoutCSV(hledgerArgs []string, filters ...Filter) (io.Reader, error) {
	args := make([]string, 0)

	args = append(args, hledgerArgs...)

	for _, f := range filters {
		if _, ok := f.(NoFilter); ok {
			continue
		}
		args = append(args, f.Build())
	}

	if h.JournalFilename != "" {
		args = append(args, "-f", h.JournalFilename)
	}

	log.Printf("running command: %s", strings.Join(args, " "))

	cmd := exec.Command(h.HledgerBinary, args...)
	result, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("error: %v", err)
		return bytes.NewBuffer(result), err
	}

	return bytes.NewBuffer(result), nil
}
