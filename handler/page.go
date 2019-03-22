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
)

func GetPage(c *gin.Context) {
	pageNum := helper.DefaultParamInt(c, "page_num", 1)

	books := dal.FindBookOrder(pageNum, consts.PerPage)

	if books == nil || len(books) == 0 {
		count := dal.CountBook()
		msg := helper.AssemblePageMsg(pageNum, int(count), consts.PerPage)
		c.HTML(http.StatusOK, consts.ErrorTmpl, msg)
		return
	}

	c.HTML(http.StatusOK, consts.PageTmpl, books)
}
