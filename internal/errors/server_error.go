package errors

type ServerError struct {
	Code int
	Msg  string
	Err  error
}

func (err ServerError) Error() string {
	return err.Msg
}

func NewServerError(msg string, err error) error {
	return ServerError{
		Code: 500,
		Msg:  msg,
		Err:  err,
	}
}

func NewServerErrorWithCode(code int, msg string, err error) error {
	return ServerError{
		Code: code,
		Msg:  msg,
		Err:  err,
	}
}
