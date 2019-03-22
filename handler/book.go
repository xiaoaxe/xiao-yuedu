// 书籍的增删改查
// author: baoqiang
// time: 2019/3/21 下午6:36
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/githubao/xiao-yuedu/dal"
	"github.com/githubao/xiao-yuedu/helper"
	"net/http"
	"github.com/githubao/xiao-yuedu/consts"
	"fmt"
)

func GetBook(c *gin.Context) {
	bid := helper.GetParamInt64(c, "book_id")

	book := dal.FindBookOne(bid)

	if book == nil {
		msg := fmt.Sprintf("book NOT FOUND: %v", bid)
		c.HTML(http.StatusNotFound, consts.ErrorTmpl, msg)
		return
	}

	c.HTML(http.StatusOK, consts.BookTmpl, book)
}
