package routers

import (
	"github.com/20gu00/k8s-dashboard/Handler"
	"github.com/20gu00/k8s-dashboard/common"
	"github.com/gin-gonic/gin"
)

//传入指针 向server注册路由(api)
func Register(engine *gin.Engine) {
	engine.Use(common.CrosMiddleware) //全局中间件

	//ping
	engine.GET("/ping", Handler.Ping) //(局部中间件,在()里边加,针对这个api)(GET("/ping",CrosMiddleware, Handler.Ping)

	//api/v1
	api := engine.Group("/api/v1") //api分组(子路由)  针对分组,api.Use()

	//nodelist
	api.GET("nodes", Handler.NodeList) ///api/v1/nodes

}
