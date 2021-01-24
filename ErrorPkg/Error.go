package ErrorPkg

// New returns an error that formats as the given text.
func New(text string) error {
	return &ErrorString{text}
}

// errorString is a trivial implementation of error.
type ErrorString struct {
	s string
}

func (e *ErrorString) Error() string {
	return e.s
}
