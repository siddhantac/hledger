package hledger

type Hledger struct {
	journalFilename string
	hledgerBinary   string
}

func New(hledgerBinary, journalFile string) Hledger {
	return Hledger{
		hledgerBinary:   hledgerBinary,
		journalFilename: journalFile,
	}
}
