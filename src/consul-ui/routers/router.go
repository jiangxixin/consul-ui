package routers

import (
	"consul-ui/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/login", &controllers.LoginController{})
	beego.Router("/api/metrics", &controllers.MetricsController{})
	beego.Router("/api/upload", &controllers.UploadController{})
}
