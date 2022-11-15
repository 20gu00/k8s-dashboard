package main

import (
	"flag"
	"github.com/20gu00/k8s-dashboard/common"
	"github.com/20gu00/k8s-dashboard/routers"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

func main() {
	//初始化日志工具(flag参数 flagset)
	klog.InitFlags(nil)
	defer klog.Flush() //存入

	//设置flagset 标准错误输出 也不输出到文件
	flag.Set("logtostderr", "false") //InitFlags->init
	flag.Set("alsologtostderr", "false")
	flag.Parse()

	//初始化
	if err := common.Init(); err != nil {
		klog.V(2).ErrorS(err, "初始化出错") //等级为二
		return                         //panic(err.Error())
	}

	//http server
	server := gin.Default() //*engine server

	//注册路由

	routers.Register(server)

	//启动server(http)

	//0.0.0.0:8888
	if err := server.Run(":8888"); err != nil {
		klog.V(2).ErrorS(err, "运行server出错")
	}
}
