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
)

func main() {
	router := gin.Default()

	db := dal.InitDb()
	defer db.Close()

	InitRouter(router)

	InitHtml(router)

	router.Run("localhost:8000")
}

func InitRouter(r *gin.Engine) {
	r.GET("/books/:bid", handler.GetBook)
	r.GET("/pages/:pid", handler.GetPage)
	r.NoRoute(handler.NotFound)
}

func InitHtml(r *gin.Engine) {
	r.LoadHTMLGlob(filepath.Join(helper.GetRootPath(), "views/*"))
}
