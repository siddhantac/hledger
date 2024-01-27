package hledger

import (
	"io"
	"strings"
)

func (h Hledger) Register(options Options) (io.Reader, error) {
	rd, err := h.execCmd("register", options)
	if err != nil {
		data, _ := io.ReadAll(rd)
		err.msg = string(strings.Trim(string(data), "\n"))
		return nil, err
	}

	return rd, nil
}
