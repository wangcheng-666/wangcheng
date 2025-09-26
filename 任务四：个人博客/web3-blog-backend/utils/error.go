package utils

type Err struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewError(code int, msg string) *Err {
	return &Err{
		Code: code,
		Msg:  msg,
	}
}

func (e *Err) Err() string {
	return e.Msg
}
