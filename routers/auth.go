package routers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego/context"
	"kabao/logs"
	"sort"
)

func auth(ctx *context.Context) {
	//
	url := ctx.Input.URL()
	fmt.Println(url)
	logs.Debug("Hello %d", 100)
	if url == "/user/login" || url == "/user/register" {
		return
	}
	var userid string
	if err := ctx.Request.ParseForm(); err != nil {
		ctx.WriteString("Error")
		return
	}
	if len(ctx.Request.Form["userid"]) <= 0 || len(ctx.Request.Form["signature"]) <= 0 || ctx.Request.Form["userid"][0] == "" {
		ctx.WriteString("Error")
		return
	}
	userid = ctx.Request.Form["userid"][0]
	fmt.Println(userid)
	sorted_keys := make([]string, 0)
	for k, _ := range ctx.Request.Form {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)
	var tmp string
	flag := 1
	for _, k := range sorted_keys {
		if k == "signature" {
			continue
		}
		if flag == 1 {
			tmp = k + "=" + ctx.Request.Form[k][0]
			flag = 0
		} else {
			tmp = tmp + "&" + k + "=" + ctx.Request.Form[k][0]
		}
	}
	h := md5.New()
	//根据用户ID获取secret
	secret := "test"
	h.Write([]byte(secret))
	sign := hex.EncodeToString(h.Sum(nil))
	tmp = tmp + "&" + sign
	h2 := md5.New()
	h2.Write([]byte(tmp))
	target := hex.EncodeToString(h2.Sum(nil))
	fmt.Println(target)
	if ctx.Request.Form["signature"][0] != target {

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
