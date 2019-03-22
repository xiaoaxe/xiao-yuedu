// 主程序
// author: baoqiang
// time: 2019/3/21 下午6:32
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/githubao/xiao-yuedu/handler"
	"github.com/githubao/xiao-yuedu/dal"
	"path/filepath"
	"github.com/githubao/xiao-yuedu/helper"
	"github.com/githubao/xiao-yuedu/tools"
)

func main() {
	router := gin.Default()

	db := dal.InitDb()
	defer db.Close()

	search := tools.InitSearch()
	defer search.Close()

	InitRouter(router)

	InitHtml(router)

	router.Run("localhost:8000")
}

func InitRouter(r *gin.Engine) {
	r.GET("/books/:book_id", handler.GetBook)
	r.GET("/pages/:page_num", handler.GetPage)
	r.GET("/search", handler.SearchBook)

	r.NoRoute(handler.NotFound)
}

func InitHtml(r *gin.Engine) {
	r.LoadHTMLGlob(filepath.Join(helper.GetRootPath(), "views/*"))
}
