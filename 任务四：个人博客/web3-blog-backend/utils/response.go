package utils

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Date any    `json:"date"`
}

func Success(data any) *Result {
	return &Result{
		Code: ok,
		Msg:  "success",
		Date: data,
	}
}

func Fail(err *Err) *Result {
	return &Result{
		Code: err.Code,
		Msg:  err.Msg,
	}
}
