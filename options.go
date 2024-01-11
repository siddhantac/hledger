package hledger

import "fmt"

type Options struct {
	account      string
	accountDepth int
	accountDrop  int
	startDate    string
	endDate      string
	outputCSV    bool
}

func NewOptions() Options { return Options{} }

func (o Options) WithAccount(account string) Options {
	o.account = account
	return o
}

func (o Options) WithAccountDepth(depth int) Options {
	o.accountDepth = depth
	return o
}

func (o Options) WithAccountDrop(drop int) Options {
	o.accountDrop = drop
	return o
}

func (o Options) WithStartDate(startDate string) Options {
	o.startDate = startDate
	return o
}

func (o Options) WithEndDate(endDate string) Options {
	o.endDate = endDate
	return o
}

func (o Options) WithOutputCSV() Options {
	o.outputCSV = true
	return o
}

func (o Options) Build() []string {
	var options []string
	if o.account != "" {
		options = append(options, "acct:"+o.account)
	}

	if o.accountDepth > 0 {
		options = append(options, fmt.Sprintf("--depth=%d", o.accountDepth))
	}

	if o.accountDrop > 0 {
		options = append(options, fmt.Sprintf("--drop=%d", o.accountDrop))
	}

	if o.startDate != "" || o.endDate != "" {
		date := fmt.Sprintf("%s..%s", o.startDate, o.endDate)
		options = append(options, "date:"+date)
	}

	if o.outputCSV {
		options = append(options, "-O", "csv")
	}
	return options
}