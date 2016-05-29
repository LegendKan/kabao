// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"kabao/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSCond(func(ctx *context.Context) bool {
			if ctx.Input.Domain() == "api.beego.me" {
				return true
			}
			return true
		}),

		beego.NSBefore(auth),

		beego.NSNamespace("/test",
			beego.NSInclude(
				&controllers.TestController{},
			)),

		beego.NSNamespace("/category",
			beego.NSInclude(
				&controllers.CategoryController{},
			),
		),

		beego.NSNamespace("/comment",
			beego.NSInclude(
				&controllers.CommentController{},
			),
		),

		beego.NSNamespace("/commoncard",
			beego.NSInclude(
				&controllers.CommoncardController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/vipcard",
			beego.NSInclude(
				&controllers.VipcardController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
