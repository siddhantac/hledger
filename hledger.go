package hledger

import (
	"io"
	"log"
)

type Hledger struct {
	journalFilename string
	hledgerBinary   string
	logger          *log.Logger
}

func New(hledgerBinary, journalFile string) Hledger {
	return Hledger{
		hledgerBinary:   hledgerBinary,
		journalFilename: journalFile,
	}
}

func NewDebug(hledgerBinary, journalFile string, logWriter io.Writer) Hledger {
	return Hledger{
		hledgerBinary:   hledgerBinary,
		journalFilename: journalFile,
		logger:          log.New(logWriter, "hledger: ", log.LstdFlags),
	}
}
