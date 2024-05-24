package errors

type ServerError struct {
	Code int
	Msg  string
	Err  error
}

func (err ServerError) Error() string {
	return err.Msg
}

func NewServerError(code int, msg string, err error) error {
	return ServerError{
		Code: code,
		Msg:  msg,
		Err:  err,
	}
}
