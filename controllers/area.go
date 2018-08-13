package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"LoveHome/models"
	//_ "github.com/astaxie/beego/cache/redis"
	//"fmt"
)

type AreaController struct {
	beego.Controller
}

func(this*AreaController)RetData(resp map[string]interface{}){
	this.Data["json"] = resp//这两句是专门用来处理json的操作，并发送到前端
	this.ServeJSON()
}

func (c *AreaController) GetArea() {
	beego.Info("connect success")

	resp := make(map[string]interface{})//通过map生成json是常用的一种方法，中括号内表示键的类型，中括号后边跟的是值的类型，但是值可能是字符串/切片/布尔/浮点等类型，所以用interface{}表示的万能类型
	resp["errno"] = models.RECODE_OK //下面两句初始化。errno键对应的值为。。。
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer c.RetData(resp)

	//从redis缓存中拿数据拿数据
/*	cache_conn, err := cache.NewCache("redis", `{"key":"lovehome","conn":":6370","dbNum":"0"}`)

	if areaData := cache_conn.Get("area");areaData!=nil{
		beego.Info("get data from cache===========")
		resp["data"] = areaData
		return
	}

*/

	//beego.Info("cache_conn.aa =",cache_conn.Get("aaa"))
	//fmt.Printf("cache_conn ,conn[aa]= %s\n",cache_conn.Get("aaa"))

	//从mysql数据库拿到area数据
	var areas  []models.Area //先定义一个变量，类型是[]models.Area，注意这个地方要是一个结构体切片，因为我们要拿所有结构体的数据。

	o := orm.NewOrm()
	num ,err :=o.QueryTable("area").All(&areas) //从指定数据库表area中读到所有的数据并存入到areas中，从表中读数据的专有函数。
//返回拿到多少条数据和错误信息
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
 	}
 	if num ==0{
		resp["errno"] = 4002
		resp["errmsg"] = "没有查到数据"
		return
	}

	resp["data"] = areas //填充areas数据到map中

	//把数据转换成json格式存入缓存
/*	json_str,err := json.Marshal(areas)
	if err != nil{
		beego.Info("encoding err")
		return
	}

	cache_conn.Put("area",json_str,time.Second*3600)

*/
		//打包成json返回给前段
	beego.Info("query data sucess ,resp =",resp,"num =",num)


}
