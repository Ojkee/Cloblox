package functools

// Special type of error wrapper, if it's initialized
// instead of typical error, it means sumulation can't start
type StrongError struct {
	consoleMessage string // On Error()
	debugMessage   string // On Debug()
}

func NewStrongError(
	consoleMessage, debugMessage string,
) *StrongError {
	return &StrongError{
		consoleMessage: consoleMessage,
		debugMessage:   debugMessage,
	}
}

func (err *StrongError) Error() string {
	return err.consoleMessage
}

func (err *StrongError) Debug() string {
	return err.debugMessage
}
