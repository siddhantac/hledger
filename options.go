package hledger

import "fmt"

type Options struct {
	account      string
	accountType  string
	accountDepth int
	accountDrop  int
	startDate    string
	endDate      string
	outputCSV    bool
	layout       LayoutType
	sortAmount   bool
	invertAmount bool
	period       PeriodType
	pretty       bool
	description  string
}

type (
	LayoutType string
	PeriodType string
)

const (
	LayoutBare      LayoutType = "bare"
	PeriodDaily     PeriodType = "--daily"
	PeriodWeekly    PeriodType = "--weekly"
	PeriodMonthly   PeriodType = "--monthly"
	PeriodQuarterly PeriodType = "--quarterly"
	PeriodYearly    PeriodType = "--yearly"
)

func NewOptions() Options { return Options{} }

func (o Options) WithDescription(description string) Options {
	o.description = description
	return o
}

func (o Options) WithPretty() Options {
	o.pretty = true
	return o
}

func (o Options) WithAccount(account string) Options {
	o.account = account
	return o
}

func (o Options) WithAccountType(accountType string) Options {
	o.accountType = accountType
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

func (o Options) WithLayout(layout LayoutType) Options {
	o.layout = layout
	return o
}

func (o Options) WithSortAmount() Options {
	o.sortAmount = true
	return o
}

func (o Options) WithInvertAmount() Options {
	o.invertAmount = true
	return o
}

func (o Options) WithPeriod(period PeriodType) Options {
	o.period = period
	return o
}

func (o Options) Build() []string {
	var options []string
	if o.account != "" {
		options = append(options, "acct:"+o.account)
	}

	if o.accountType != "" {
		options = append(options, "type:"+o.accountType)
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

	if o.layout != "" {
		options = append(options, "--layout", string(o.layout))
	}

	if o.sortAmount {
		options = append(options, "--sort-amount")
	}

	if o.invertAmount {
		options = append(options, "--invert")
	}

	if o.period != "" {
		options = append(options, string(o.period))
	}

	if o.outputCSV {
		options = append(options, "-O", "csv")
	}

	if o.pretty {
		options = append(options, "--pretty")
	}

	if o.description != "" {
		options = append(options, "desc:\""+o.description+"\"")
	}
	return options
}
