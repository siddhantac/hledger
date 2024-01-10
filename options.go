package hledger

import "fmt"

type Options struct {
	account      string
	accountDepth int
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

func (o Options) Build() []string {
	var options []string
	if o.account != "" {
		options = append(options, "acct:"+o.account)
	}

	if o.accountDepth > 0 {
		options = append(options, fmt.Sprintf("--depth=%d", o.accountDepth))
	}
	return options
}
