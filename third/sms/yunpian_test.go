package sms

import (
	"testing"
)

func TestSendSingleSMS(t *testing.T) {
	err := SendSingleSMS("15991684027", "【我的卡包】您的短信验证码是520520，欢迎注册“我的卡包”账号。")
	if err == nil {
		t.Log("Send Successfully")
	} else {
		t.Error("Send Error: " + err.Error())
	}
}
