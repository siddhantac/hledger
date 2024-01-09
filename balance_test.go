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
	rd, err := hl.Balance()
	assert.NoError(t, err)
	b, _ := io.ReadAll(rd)
	fmt.Println(string(b))
}
