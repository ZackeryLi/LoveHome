package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"LoveHome/models"

)


type UserController struct {
	beego.Controller
}

func(this*UserController)RetData(resp map[string]interface{}){
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this*UserController)Reg(){
	resp := make(map[string]interface{}) //定义一个获得结构体数据的变量
	defer this.RetData(resp)

//获取前端传过来的json数据
	json.Unmarshal(this.Ctx.Input.RequestBody,&resp)//解包json数据，参数：将解包的json数据存放到结构体中
/*
	mobile:"111"
	password:"111"
	sms_code:"111"

	beego.Info(`resp["mobile"] =`,resp["mobile"])
	beego.Info(`resp["password"] =`,resp["password"])
	beego.Info(`resp["sms_code"] =`,resp["sms_code"])
*/

//将注册信息插入数据库
o:=orm.NewOrm()
user := models.User{}
user.Password_hash = resp["password"].(string) //因为我们的字典中的value是万能类型的，所以此处要进行转换为string
user.Name = resp["mobile"].(string) //由于网页端传过来的数据只有手机号和密码（可以在浏览器端查看），所以就用手机号当作用户名。
user.Mobile = resp["mobile"].(string)
id,err:=o.Insert(&user)
if err != nil{
	resp["errno"] = 4002
	resp["errmsg"] = "注册失败"
	return
}

beego.Info("reg success ,id = ",id)
	resp["errno"] = 0
	resp["errmsg"] = "注册成功"

	this.SetSession("name",user.Name) //在注册结束后将需要的信息添加进session，后面网页会调用GetSessionData函数来拿到session的内容。

}
/*
func (this*UserController)Postavatar(){

	resp := make(map[string]interface{})
	defer this.RetData(resp)
	//1.获取上传的一个文件
	fileData,hd,err := this.GetFile("avatar")
	if err != nil {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		beego.Info("===========11111")
		return
	}
	//2.得到文件后缀
	suffix := path.Ext(hd.Filename) //a.jpg.avi


	//3.存储文件到fastdfs上
	fdfsClient,err := fdfs_client.NewFdfsClient("conf/client.conf")
	if err != nil{
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		beego.Info("===========22222222")

		return
	}
	fileBuffer := make([]byte, hd.Size)
	_, err = fileData.Read(fileBuffer)
	if err != nil {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		beego.Info("===========33333")

		return
	}
	DataResponse, err := fdfsClient.UploadByBuffer(fileBuffer, suffix[1:])//aa.jpg

	if err != nil {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		beego.Info("===========44444")

		return
	}

	//DataResponse.GroupName
	//DataResponse.RemoteFileId   //group/mm/00/00231312313131231.jpg

	//4.从session里拿到user_id
	user_id := this.GetSession("user_id")
	var user models.User
	//5.更新用户数据库中的内容
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	qs.Filter("Id",user_id).One(&user)
	user.Avatar_url = DataResponse.RemoteFileId

	_,errUpdate := o.Update(&user)
	if errUpdate != nil{
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}

	urlMap:= make(map[string]string)
	//Avaurl := "192.168.152.138:8899"+DataResponse.RemoteFileId
	urlMap["avatar_url"] = "http://192.168.152.138:8899/"+DataResponse.RemoteFileId
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = urlMap

}
*/