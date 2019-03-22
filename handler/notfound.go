// 404
// author: baoqiang
// time: 2019/3/21 下午9:41
package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/githubao/xiao-yuedu/consts"
)

func NotFound(c *gin.Context) {
	path := c.Request.URL.Path
	msg := fmt.Sprintf("path NOT FOUND: %v", path)
	c.HTML(http.StatusNotFound, consts.ErrorTmpl, msg)
}
