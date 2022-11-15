package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//json编码gin.H响应
func RespErr(c *gin.Context, msg string) {
	//响应成功,使用自定义的信息传递实际响应的信息(响应头)
	c.JSON(http.StatusOK, gin.H{ //map[string]any
		"code": 1,
		"msg":  msg,
	})
}

func RespOk(c *gin.Context, data interface{}) {
	resp, ok := data.(gin.H) //接口断言,成与不成,成功有值类型gin.H,不成功零值类型interface{}
	if !ok {                 //断言失败
		resp = gin.H{} //空接口赋值,空
		resp["data"] = data
	}

	resp["code"] = 0
	resp["msg"] = "success"
	c.JSON(http.StatusOK, resp)
}
