package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myblogweb/models"
	"myblogweb/utils"
	"time"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	this.TplName = "register.html"
}

//处理注册
func (this *RegisterController) Post() {
	//获取表单信息
	un := this.GetString("username")
	pd := this.GetString("password")
	rePassword := this.GetString("rePassword")
	fmt.Println(un, pd , rePassword)

	//注册之前先判断该用户名是否已经被注册，如果已经注册，返回错误
	id := models.QueryUserWithUsername(un)
	fmt.Println("id:", id)
	if id > 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "用户名已经存在"}
		this.ServeJSON()
		return
	}

	//注册用户名和密码
	//存储的密码是md5后的数据，那么在登录的验证的时候，也是需要将用户的密码md5之后和数据库里面的密码进行判断
	pd = utils.MD5(pd)
	fmt.Println("md5后：", pd)
	var user = models.User{0, un, pd, 0, time.Now().Unix()}
	_, err := models.InsertUser(user)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "注册失败"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "注册成功"}
	}
	this.ServeJSON()
}
