package hledger

import (
	"io"
)

func (h Hledger) IncomeStatement(options Options) (io.Reader, error) {
	rd, err := h.execCmd("incomestatement", options)
	if err != nil {
		data, _ := io.ReadAll(rd)
		err.msg = string(data)
		return nil, err
	}

	return rd, nil
}
