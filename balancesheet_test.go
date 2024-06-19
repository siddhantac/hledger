package hledger_test

import (
	"io"
	"testing"

	"github.com/siddhantac/hledger"
	"github.com/stretchr/testify/assert"
)

func TestBalanceSheet(t *testing.T) {
	hl := hledger.New("hledger", "data/hledger.journal")

	opts := hledger.NewOptions()
	rd, err := hl.BalanceSheet(opts)
	assert.NoError(t, err)

	expected := `Balance Sheet 2022-05-17

                                          ||  2022-05-17 
==========================================++=============
 Assets                                   ||             
------------------------------------------++-------------
 assets:bank:chase                        ||   $5,348.35 
 assets:bank:maybank                      || $-12,300.00 
 assets:bank:stanchart                    ||      $22.59 
 assets:invest:crypto                     ||   $2,000.00 
 assets:invest:stocks                     ||   $1,000.00 
------------------------------------------++-------------
                                          ||  $-3,929.06 
==========================================++=============
 Liabilities                              ||             
------------------------------------------++-------------
 liabilities:credit_card:american_express ||    $-514.86 
 liabilities:credit_card:citibank         ||     $-36.11 
------------------------------------------++-------------
                                          ||    $-550.97 
==========================================++=============
 Net:                                     ||  $-3,378.09 
`

	got, err := io.ReadAll(rd)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(got))
}

func TestBalanceSheetWithTree(t *testing.T) {
	hl := hledger.New("hledger", "data/hledger.journal")

	opts := hledger.NewOptions().WithTree()
	rd, err := hl.BalanceSheet(opts)
	assert.NoError(t, err)

	expected := `Balance Sheet 2022-05-17

                         ||  2022-05-17 
=========================++=============
 Assets                  ||             
-------------------------++-------------
 assets                  ||  $-3,929.06 
   bank                  ||  $-6,929.06 
     chase               ||   $5,348.35 
     maybank             || $-12,300.00 
     stanchart           ||      $22.59 
   invest                ||   $3,000.00 
     crypto              ||   $2,000.00 
     stocks              ||   $1,000.00 
-------------------------++-------------
                         ||  $-3,929.06 
=========================++=============
 Liabilities             ||             
-------------------------++-------------
 liabilities:credit_card ||    $-550.97 
   american_express      ||    $-514.86 
   citibank              ||     $-36.11 
-------------------------++-------------
                         ||    $-550.97 
=========================++=============
 Net:                    ||  $-3,378.09 
`

	got, err := io.ReadAll(rd)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(got))
}
