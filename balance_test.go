package hledger_test

import (
	"testing"

	"github.com/siddhantac/hledger"
	"github.com/stretchr/testify/assert"
)

func TestBalance(t *testing.T) {
	hl := hledger.New("hledger", "data/hledger.journal")

	opts := hledger.NewOptions().WithAccount("maybank")
	rd, err := hl.Balance(opts.WithOutputCSV())
	assert.NoError(t, err)
	records, err := hledger.ParseCSV(rd)
	assert.NoError(t, err)

	expected := [][]string{
		{"account", "balance"},
		{"assets:bank:maybank", "$-12300.00"},
		{"total", "$-12300.00"},
	}
	assert.Equal(t, expected, records)
}
