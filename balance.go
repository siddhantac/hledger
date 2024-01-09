package hledger

import (
	"fmt"
	"io"
)

func (h Hledger) Balance( /*filters ...Filter*/ ) (io.Reader, error) {
	rd, err := h.execWithoutCSV([]string{"balance"}) //, filters...)
	if err != nil {
		data, _ := io.ReadAll(rd)
		return nil, fmt.Errorf(string(data))
	}

	return rd, nil
}
