package routers

import (
	"fmt"
	"github.com/astaxie/beego/context"
	"kabao/logs"
)

func init() {
	//log:=logs.GetLogger()
}

func auth(ctx *context.Context) {
	//
	url := ctx.Input.URL()
	fmt.Println(url)
	logs.GetLogger().Debug("Hello")
	if url == "/user/login" || url == "/user/register" {
		return
	}
	//auth := ctx.Request.URL.Query().Get("auth")
	/*
		if len(auth) <= 0 || !controllers.AuthMaster(auth) {
			ret := common.BuildResult(common.UnAuth, "")
			//ctx.Redirect(302, "/login")
			b, _ := json.Marshal(ret)
			ctx.WriteString(string(b))
		}
	*/
}
