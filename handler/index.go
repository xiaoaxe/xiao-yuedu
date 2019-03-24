// index
// author: baoqiang
// time: 2019/3/24 下午7:31
package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/githubao/xiao-yuedu/consts"
)

func Index(c *gin.Context){
	c.HTML(http.StatusOK,consts.IndexTmpl,nil)
}