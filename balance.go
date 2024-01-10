package hledger

import (
	"fmt"
	"io"
)

func (h Hledger) Balance(options Options) (io.Reader, error) {
	args := []string{"balance"}
	args = append(args, options.Build()...)
	rd, err := h.execWithoutCSV(args)
	if err != nil {
		data, _ := io.ReadAll(rd)
		return nil, fmt.Errorf(string(data))
	}

	return rd, nil
}
