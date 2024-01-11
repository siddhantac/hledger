package hledger

type Error struct {
	err error
	msg string
}

func (e *Error) Error() string {
	return e.err.Error()
}
func (e *Error) Msg() string {
	return e.msg
}
