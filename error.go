package hledger

type Error struct {
	err     error
	msg     string
	cmdArgs []string
}

func (e *Error) Error() string {
	return e.err.Error()
}
func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Args() []string { return e.cmdArgs }
