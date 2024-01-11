package hledger_test

import (
	"testing"

	"github.com/siddhantac/hledger"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	hl := hledger.New("hledger", "data/hledger.journal")

	opts := hledger.NewOptions().WithAccount("maybank")
	rd, err := hl.Register(opts.WithOutputCSV())
	assert.NoError(t, err)
	records, err := hledger.ParseCSV(rd)
	assert.NoError(t, err)

	expected := [][]string{
		{"txnidx", "date", "code", "description", "account", "amount", "total"},
		{"2", "2021-12-08", "", "Rent", "assets:bank:maybank", "$-2300.00", "$-2300.00"},
		{"6", "2022-01-05", "", "Moving it around", "assets:bank:maybank", "$-10000.00", "$-12300.00"},
	}
	assert.Equal(t, expected, records)
}
