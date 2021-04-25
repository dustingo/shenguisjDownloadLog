package controllers

import (
	"shenguisjDownloadLog/models"

	"log"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about dl
type DownloadGameController struct {
	beego.Controller
}

// @Title Get
// @Description put the time
// @Param   date    query    string  true        "The date you want to download"
// @Param   user    query    string  true        "The user you got from dou"
// @Param   auth    query    string  true        "The auth you got from dou"
// @Success 200 {int} 200
// @Failure 403 date is nil
// @router / [get]
func (D *DownloadGameController) Get() {
	log.Printf("REMOTEIP:%s RequestURL:%s\n", D.Ctx.Request.RemoteAddr, D.Ctx.Request.RequestURI)
	timeStr := D.GetString("date")
	user := D.GetString("user")
	auth := D.GetString("auth")
	authInfo := models.Admin{}
	if user == authInfo.GetUser() && auth == authInfo.GetAuth() {
		err := models.DownloadGameLog(timeStr)
		if err != nil {
			D.Data["json"] = err.Error()
		} else {
			D.Data["json"] = "ok"
		}

	} else {
		err := "user or auth wrong"
		D.Data["json"] = err
	}
	D.ServeJSON()
}

// @Title DownloadFile
// @Description download log file
// @Param   user    query    string  true        "The user you got from dou"
// @Param   auth    query    string  true        "The auth you got from dou"
// @Success 200 [string] download ok
// @Failure 403 file is not existed
// @router /download [get]
func (D *DownloadGameController) DownloadFile() {
	user := D.GetString("user")
	auth := D.GetString("auth")
	authInfo := models.Admin{}
	if user == authInfo.GetUser() && auth == authInfo.GetAuth() {
		file := models.DownloadComprocessFile()
		D.Ctx.Output.Download(file)
		D.Data["json"] = "download ok"
	} else {
		D.Data["json"] = "user or pass error"
	}
	D.ServeJSON()
}
