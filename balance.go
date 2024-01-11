package hledger

import (
	"io"
)

func (h Hledger) Balance(options Options) (io.Reader, error) {
	rd, err := h.execCmd("balance", options)
	if err != nil {
		e := &Error{err: err}
		data, _ := io.ReadAll(rd)
		e.msg = string(data)
		return nil, e
	}

	return rd, nil
}
