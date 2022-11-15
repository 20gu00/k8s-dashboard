package routers

import (
	"github.com/20gu00/k8s-dashboard/Handler"
	"github.com/gin-gonic/gin"
)

//传入指针
//向server注册路由(api)
func Register(engine *gin.Engine) {
	//ping
	engine.GET("/ping", Handler.Ping) //get /ping handlerfunc_name(controller)

	//api/v1
	//api分组(子路由)
	api := engine.Group("/api/v1")

	//nodelist
	///api/v1/nodes
	api.GET("nodes", Handler.NodeList)

}
