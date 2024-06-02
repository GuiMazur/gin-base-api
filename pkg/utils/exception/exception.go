package exception

type Exception struct {
	message string
	code    int
	error   string
}

func New(message string, code int, error string) *Exception {
	return &Exception{
		message: message,
		code:    code,
		error:   error,
	}
}

func (e *Exception) Error() string {
	return e.error
}
