package routers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego/context"
	"kabao/logs"
	"kabao/models"
	"sort"
	"strconv"
)

func auth(ctx *context.Context) {
	//
	url := ctx.Input.URL()
	fmt.Println(url)
	logs.Debug("Hello %s", url)
	if url == "/v1/user/login" || url == "/v1/user/register" || url == "/v1/test/" {
		return
	}
	var userid string
	if err := ctx.Request.ParseForm(); err != nil {
		ctx.WriteString("Error")
		return
	}
	var tokenid int
	if len(ctx.Request.Form["userid"]) <= 0 || len(ctx.Request.Form["signature"]) <= 0 || ctx.Request.Form["userid"][0] == "" || len(ctx.Request.Form["tokenid"][0]) <= 0 {
		ctx.WriteString("Error")
		return
	}
	userid = ctx.Request.Form["userid"][0]
	tokenid, err := strconv.Atoi(ctx.Input.Param("tokenid"))
	if err != nil {
		ctx.WriteString("Error")
		return
	}
	//根据token ID获取secret
	secret, err := models.GetToken(tokenid)
	if err != nil {
		ctx.WriteString("Error")
		return
	}
	fmt.Println(userid)
	sorted_keys := make([]string, 0)
	params := ctx.Input.Params()
	for k, _ := range params { //ctx.Request.Form
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
			tmp = k + "=" + params[k]
			flag = 0
		} else {
			tmp = tmp + "&" + k + "=" + params[k]
		}
	}
	h := md5.New()
	h.Write([]byte(secret))
	sign := hex.EncodeToString(h.Sum(nil))
	tmp = tmp + "&" + sign
	h2 := md5.New()
	h2.Write([]byte(tmp))
	target := hex.EncodeToString(h2.Sum(nil))
	fmt.Println(target)
	if params["signature"] != target {

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
