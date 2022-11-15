package Handler

import (
	"github.com/20gu00/k8s-dashboard/common"
	"github.com/20gu00/k8s-dashboard/k8sdo"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

func NodeList(ctx *gin.Context) {
	nodeList, err := k8sdo.GetNode(common.NewKubeClient(), "")
	if err != nil {
		klog.V(2).ErrorS(err, "获取node列表出错", "Handler", "NodeList")
		common.RespErr(ctx, err.Error())
		return
	}
	common.RespOk(ctx, gin.H{"nodes:": nodeList})
}
