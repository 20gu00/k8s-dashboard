package Handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//handlerfunc
func Ping(ctx *gin.Context) {
	//responese(json)
	ctx.JSONP(http.StatusOK, "pong") //code interface{}

}
