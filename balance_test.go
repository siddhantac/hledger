package hledger_test

import (
	"encoding/csv"
	"errors"
	"hledger-go"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBalance(t *testing.T) {
	hl := hledger.New("hledger", "data/hledger.journal")

	opts := hledger.NewOptions().WithAccount("maybank")
	rd, err := hl.Balance(opts)
	assert.NoError(t, err)
	b, _ := io.ReadAll(rd)
	assert.NotEmpty(t, b)

	rd, err = hl.Balance(opts.WithOutputCSV())
	assert.NoError(t, err)
	records := parseCSV(t, rd)

	expected := [][]string{
		{"account", "balance"},
		{"assets:bank:maybank", "$-12300.00"},
		{"total", "$-12300.00"},
	}
	assert.Equal(t, expected, records)
}

func parseCSV(t *testing.T, r io.Reader) [][]string {
	t.Helper()
	result := make([][]string, 0)
	csvrdr := csv.NewReader(r)
	for {
		rec, err := csvrdr.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		result = append(result, rec)
	}
	return result
}
