package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/siddhantac/hledger"
)

func main() {
	var journalFile, hledgerExecutable, desc string
	var isDebug bool
	flag.StringVar(&journalFile, "file", "data/hledger.journal", "journal filename")
	flag.StringVar(&hledgerExecutable, "exe", "hledger", "hledger executable")
	flag.StringVar(&desc, "desc", "", "account filter")
	flag.BoolVar(&isDebug, "debug", false, "run in debug mode")
	flag.Parse()

	var logfile *os.File
	if isDebug {
		logfilename := "debug.log"
		var err error
		logfile, err = os.OpenFile(logfilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer logfile.Close()
	}

	hl := hledger.NewDebug(hledgerExecutable, journalFile, logfile)

	options := hledger.NewOptions().WithOutputCSV()
	if desc != "" {
		options = options.WithDescription(desc)
	}

	fmt.Println(options.Build())

	rd, err := hl.Register(options)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := io.ReadAll(rd)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}

// func parseCSV(r io.Reader) ([][]string, error) {
// 	result := make([][]string, 0)
// 	csvrdr := csv.NewReader(r)
// 	// csvrdr.Read() // skip 1 line
// 	for {
// 		rec, err := csvrdr.Read()
// 		if errors.Is(err, io.EOF) {
// 			break
// 		}
// 		if err != nil {
// 			return nil, err
// 		}
// 		result = append(result, rec)
// 	}
// 	return result, nil
// }
