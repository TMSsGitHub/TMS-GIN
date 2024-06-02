package resp

type R struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

const RES = "res"

func Success(data any) R {
	return R{
		Code: 200,
		Msg:  "success",
		Data: data,
	}
}

func Fail(msg string) R {
	return R{
		Code: 500,
		Msg:  msg,
		Data: nil,
	}
}

func FailWithCode(code int, msg string) R {
	return R{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
