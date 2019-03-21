// 处理分页数据
// author: baoqiang
// time: 2019/3/21 下午7:44
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/githubao/xiao-yuedu/dal"
	"net/http"
	"github.com/githubao/xiao-yuedu/helper"
	"github.com/githubao/xiao-yuedu/consts"
	"fmt"
)

func GetPage(c *gin.Context) {
	//TODO 实现分页查询
	pid := helper.GetParamInt64(c, "pid")

	books := dal.FindOrder(pid, consts.PerPage)

	if books == nil {
		msg := fmt.Sprintf("page NOT FOUND: %v", pid)
		c.JSON(http.StatusOK, helper.NotFound(msg))
		return
	}

	c.HTML(http.StatusOK, consts.PageTmpl, books)
}
