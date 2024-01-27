package hledger

import (
	"io"
)

func (h Hledger) Balance(options Options) (io.Reader, error) {
	rd, err := h.execCmd("balance", options)
	if err != nil {
		data, _ := io.ReadAll(rd)
		err.msg = string(data)
		return nil, err
	}

	return rd, nil
}

func (h Hledger) Assets(options Options) (io.Reader, error) {
	options = options.WithAccountType("a")
	rd, err := h.execCmd("balance", options)
	if err != nil {
		data, _ := io.ReadAll(rd)
		err.msg = string(data)
		return nil, err
	}

	return rd, nil
}

func (h Hledger) Expenses(options Options) (io.Reader, error) {
	options = options.WithAccountType("x")
	rd, err := h.execCmd("balance", options)
	if err != nil {
		data, _ := io.ReadAll(rd)
		err.msg = string(data)
		return nil, err
	}

	return rd, nil
}

func (h Hledger) Liabilities(options Options) (io.Reader, error) {
	options = options.WithAccountType("l")
	rd, err := h.execCmd("balance", options)
	if err != nil {
		data, _ := io.ReadAll(rd)
		err.msg = string(data)
		return nil, err
	}

	return rd, nil
}

func (h Hledger) Revenue(options Options) (io.Reader, error) {
	options = options.WithAccountType("r")
	rd, err := h.execCmd("balance", options)
	if err != nil {
		data, _ := io.ReadAll(rd)
		err.msg = string(data)
		return nil, err
	}

	return rd, nil
}
