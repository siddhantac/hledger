package hledger_test

import (
	"hledger-go"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptions(t *testing.T) {
	tests := map[string]struct {
		opts     func() hledger.Options
		expected []string
	}{
		"with account": {
			opts: func() hledger.Options {
				return hledger.NewOptions().WithAccount("dbs")
			},
			expected: []string{"acct:dbs"},
		},
		"with account depth": {
			opts: func() hledger.Options {
				return hledger.NewOptions().WithAccountDepth(1)
			},
			expected: []string{"--depth=1"},
		},
		"with account drop": {
			opts: func() hledger.Options {
				return hledger.NewOptions().WithAccountDrop(1)
			},
			expected: []string{"--drop=1"},
		},
		"with start date": {
			opts: func() hledger.Options {
				return hledger.NewOptions().WithStartDate("2020-01-01")
			},
			expected: []string{"date:2020-01-01.."},
		},
		"with end date": {
			opts: func() hledger.Options {
				return hledger.NewOptions().WithEndDate("2020-02-01")
			},
			expected: []string{"date:..2020-02-01"},
		},
		"with start and end date": {
			opts: func() hledger.Options {
				return hledger.NewOptions().WithStartDate("2020-01-01").WithEndDate("2020-02-01")
			},
			expected: []string{"date:2020-01-01..2020-02-01"},
		},
		"with output csv": {
			opts: func() hledger.Options {
				return hledger.NewOptions().WithOutputCSV()
			},
			expected: []string{"-O", "csv"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.opts().Build())
		})
	}
}