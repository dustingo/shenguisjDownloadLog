package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["shenguisjDownloadLog/controllers:DownloadGameController"] = append(beego.GlobalControllerRouter["shenguisjDownloadLog/controllers:DownloadGameController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["shenguisjDownloadLog/controllers:DownloadGameController"] = append(beego.GlobalControllerRouter["shenguisjDownloadLog/controllers:DownloadGameController"],
        beego.ControllerComments{
            Method: "DownloadFile",
            Router: "/download",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
