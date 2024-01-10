package hledger

type Options struct {
	account string
}

func NewOptions() Options { return Options{} }

func (o Options) WithAccount(account string) Options {
	o.account = account
	return o
}

func (o Options) Build() []string {
	var options []string
	if o.account != "" {
		options = append(options, "acct:"+o.account)
	}
	return options
}
