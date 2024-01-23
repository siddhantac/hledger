package hledger

import (
	"io"
)

func (h Hledger) BalanceSheet(options Options) (io.Reader, error) {
	rd, err := h.execCmd("balancesheet", options)
	if err != nil {
		e := &Error{err: err}
		data, _ := io.ReadAll(rd)
		e.msg = string(data)
		return nil, e
	}

	return rd, nil
}
