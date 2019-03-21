// 404
// author: baoqiang
// time: 2019/3/21 下午9:41
package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/githubao/xiao-yuedu/helper"
	"fmt"
)

func NotFound(c *gin.Context) {
	path := c.Request.URL.Path
	msg := fmt.Sprintf("path NOT FOUND: %v", path)
	c.JSON(http.StatusOK, helper.NotFound(msg))
}
