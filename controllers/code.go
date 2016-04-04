package controller

const (
	OK=iota   //正常，无错误
	SomeError	//有错误
	PhoneHaveRegistered
	LoginError
	
)

var CodeMsgs = map[int]string(
	OK:"OK",
	SomeError:"Have Some Errors",
	PhoneHaveRegistered:"The Phone Has been Registered",
	LoginError:"Username or Password is Incorrect",
)

//根据返回码获得文字说明
func GetCodeMsg(code int) string {
	if v, ok:=CodeMsgs[code];ok{
		return v
	}else{
		return "Unknown"
	}
}

func AddCode(code int, msg string) bool {
	if v, ok:CodeMsgs[code];ok{
		return false
	}else{
		CodeMsgs[code]=msg
		return true
	}
}