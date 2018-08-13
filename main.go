package main

import (
	_ "LoveHome/routers"
	"github.com/astaxie/beego"
	"strings"
	"net/http"
	"github.com/astaxie/beego/context"
	_ "LoveHome/models"
)

func main() {
	ignoreStaticPath()
	//models.TestUploadByFilename("main.go")
	beego.BConfig.WebConfig.Session.SessionOn = true //beego中要实现session的话，就需要这条语句把开关打开，或者在配置文件app.conf中写上session=true
	beego.Run(":8899") //此处可以换端口，就是把端口换成了8899
}
//下面两个函数就是url重定向
func ignoreStaticPath() {

	//透明static
	beego.SetStaticPath("group1/M00/","fdfs/storage_data/data/")
//所有路由为“/”和“/*”的操作都会进入到TransparentStatic函数
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	beego.Debug("request url: ", orpath)
	//如果请求uri还有api字段,说明是指令应该取消静态资源路径重定向
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+ctx.Request.URL.Path)
}