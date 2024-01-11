package hledger

import (
	"io"
)

func (h Hledger) Register(options Options) (io.Reader, error) {
	rd, err := h.execCmd("register", options)
	if err != nil {
		e := &Error{err: err}
		data, _ := io.ReadAll(rd)
		e.msg = string(data)
		return nil, e
	}

	return rd, nil
}
