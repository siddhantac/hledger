package hledger

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBalance(t *testing.T) {
	hl := Hledger{
		HledgerBinary: "hledger",
	}

	opts := NewOptions().WithAccount("dbs").WithAccountDepth(1)
	rd, err := hl.Balance(opts)
	assert.NoError(t, err)
	b, _ := io.ReadAll(rd)
	fmt.Println(string(b))
}
