package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["kabao/controllers:CategoryController"] = append(beego.GlobalControllerRouter["kabao/controllers:CategoryController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:CategoryController"] = append(beego.GlobalControllerRouter["kabao/controllers:CategoryController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:CategoryController"] = append(beego.GlobalControllerRouter["kabao/controllers:CategoryController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:CategoryController"] = append(beego.GlobalControllerRouter["kabao/controllers:CategoryController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:CategoryController"] = append(beego.GlobalControllerRouter["kabao/controllers:CategoryController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:CommentController"] = append(beego.GlobalControllerRouter["kabao/controllers:CommentController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:CommentController"] = append(beego.GlobalControllerRouter["kabao/controllers:CommentController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:CommentController"] = append(beego.GlobalControllerRouter["kabao/controllers:CommentController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:CommentController"] = append(beego.GlobalControllerRouter["kabao/controllers:CommentController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:CommentController"] = append(beego.GlobalControllerRouter["kabao/controllers:CommentController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:CommoncardController"] = append(beego.GlobalControllerRouter["kabao/controllers:CommoncardController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:CommoncardController"] = append(beego.GlobalControllerRouter["kabao/controllers:CommoncardController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:CommoncardController"] = append(beego.GlobalControllerRouter["kabao/controllers:CommoncardController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:CommoncardController"] = append(beego.GlobalControllerRouter["kabao/controllers:CommoncardController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:CommoncardController"] = append(beego.GlobalControllerRouter["kabao/controllers:CommoncardController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:TestController"] = append(beego.GlobalControllerRouter["kabao/controllers:TestController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:TestController"] = append(beego.GlobalControllerRouter["kabao/controllers:TestController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:TokenController"] = append(beego.GlobalControllerRouter["kabao/controllers:TokenController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:TokenController"] = append(beego.GlobalControllerRouter["kabao/controllers:TokenController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:TokenController"] = append(beego.GlobalControllerRouter["kabao/controllers:TokenController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:TokenController"] = append(beego.GlobalControllerRouter["kabao/controllers:TokenController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:TokenController"] = append(beego.GlobalControllerRouter["kabao/controllers:TokenController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:UserController"] = append(beego.GlobalControllerRouter["kabao/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:UserController"] = append(beego.GlobalControllerRouter["kabao/controllers:UserController"],
		beego.ControllerComments{
			"SignUp",
			`/signup`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:UserController"] = append(beego.GlobalControllerRouter["kabao/controllers:UserController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:UserController"] = append(beego.GlobalControllerRouter["kabao/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:UserController"] = append(beego.GlobalControllerRouter["kabao/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:UserController"] = append(beego.GlobalControllerRouter["kabao/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:VipcardController"] = append(beego.GlobalControllerRouter["kabao/controllers:VipcardController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:VipcardController"] = append(beego.GlobalControllerRouter["kabao/controllers:VipcardController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:VipcardController"] = append(beego.GlobalControllerRouter["kabao/controllers:VipcardController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:VipcardController"] = append(beego.GlobalControllerRouter["kabao/controllers:VipcardController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["kabao/controllers:VipcardController"] = append(beego.GlobalControllerRouter["kabao/controllers:VipcardController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
