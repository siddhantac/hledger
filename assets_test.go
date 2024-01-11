package hledger_test

import (
	"hledger-go"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssets(t *testing.T) {
	hl := hledger.New("hledger", "data/hledger.journal")

	opts := hledger.NewOptions()
	rd, err := hl.Assets(opts.WithOutputCSV())
	assert.NoError(t, err)
	records := parseCSV(t, rd)

	expected := [][]string{
		{"account", "balance"},
		{"assets:bank:chase", "$5348.35"},
		{"assets:bank:maybank", "$-12300.00"},
		{"assets:bank:stanchart", "$22.59"},
		{"assets:invest:crypto", "$2000.00"},
		{"assets:invest:stocks", "$1000.00"},
		{"total", "$-3929.06"},
	}
	assert.Equal(t, expected, records)
}
