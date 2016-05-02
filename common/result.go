package common

type Result struct {
	ResultCode int
	Msg        string
	Data       interface{}
}

func BuildResult(code int, data interface{}) *Result {
	msg := GetCodeMsg(code)
	return &Result{code, msg, data}
}
