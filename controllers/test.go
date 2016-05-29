package controllers

import (
	"github.com/astaxie/beego"
)

// oprations for User
type TestController struct {
	beego.Controller
}

func (c *TestController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

// @Title Post
// @Description create User
// @Success 201 {int} models.User
// @Failure 403 body is empty
// @router / [post]
func (c *TestController) Post() {
	r := BuildResult(OK, "Test: Post OK")
	c.Data["json"] = *r
	c.ServeJSON()
}

// @Title Get
// @Description get User by id
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router / [get]
func (c *TestController) Get() {
	r := BuildResult(OK, "Test: Get OK")
	c.Data["json"] = *r
	c.ServeJSON()
}
