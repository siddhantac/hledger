package hledger_test

import (
	"hledger-go"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	hl := hledger.New("hledger", "data/hledger.journal")

	opts := hledger.NewOptions().WithAccount("maybank")
	rd, err := hl.Register(opts)
	assert.NoError(t, err)
	b, _ := io.ReadAll(rd)
	assert.NotEmpty(t, b)

	rd, err = hl.Register(opts.WithOutputCSV())
	assert.NoError(t, err)
	records := parseCSV(t, rd)

	expected := [][]string{
		{"txnidx", "date", "code", "description", "account", "amount", "total"},
		{"2", "2021-12-08", "", "Rent", "assets:bank:maybank", "$-2300.00", "$-2300.00"},
		{"6", "2022-01-05", "", "Moving it around", "assets:bank:maybank", "$-10000.00", "$-12300.00"},
	}
	assert.Equal(t, expected, records)
}
