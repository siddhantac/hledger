package hledger

import (
	"io"
)

func (h Hledger) BalanceSheet(options Options) (io.Reader, error) {
	rd, err := h.execCmd("balancesheet", options)
	if err != nil {
		data, _ := io.ReadAll(rd)
		err.msg = string(data)
		return nil, err
	}

	return rd, nil
}
