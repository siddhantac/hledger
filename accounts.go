package hledger

import "io"

func (h Hledger) Accounts(options Options) (io.Reader, error) {
	rd, err := h.execCmd("accounts", options)
	if err != nil {
		data, _ := io.ReadAll(rd)
		err.msg = string(data)
		return nil, err
	}

	return rd, nil
}
