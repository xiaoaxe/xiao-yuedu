// 热门排序
// author: baoqiang
// time: 2019/3/24 下午7:39
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/githubao/xiao-yuedu/dal"
	"net/http"
	"github.com/githubao/xiao-yuedu/consts"
	"github.com/githubao/xiao-yuedu/helper"
	"github.com/githubao/xiao-yuedu/models"
)

func GetHot(c *gin.Context) {
	pageNum := helper.DefaultQueryInt(c, "p", 1)

	books := dal.FindHot(pageNum, consts.PerPage)
	count := dal.CountBook()

	if len(books) == 0 {
		msg := helper.AssemblePageMsg(pageNum, int(count), consts.PerPage)
		c.HTML(http.StatusOK, consts.ErrorTmpl, msg)
		return
	}

	c.HTML(http.StatusOK, consts.HotTmpl, models.Hot{
		Books:    books,
		PageTmpl: BuildPaginateHtml(c.Request.URL.String(), pageNum, int(count), consts.PerPage),
	})
}
