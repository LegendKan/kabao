package controllers

type Result struct {
	ResultCode int
	Msg string
	Data interface
}

func BuildResult(code int, msg string, data interface) Result {
	return &Result{code,msg,data}
}
