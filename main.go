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
	"html/template"
)

func main() {
	router := gin.Default()

	db := dal.InitDb()
	defer db.Close()

	if helper.SearchEnabled() {
		search := tools.InitSearch()
		defer search.Close()
	}

	InitRouter(router)

	InitHtml(router)

	router.Run("localhost:8000")
	//router.Run(":8000")
}

func InitRouter(r *gin.Engine) {
	r.GET("/", handler.Index)
	r.GET("/books/:book_id", handler.GetBook)
	r.GET("/pages/:page_num", handler.GetPage)
	r.GET("/download", handler.DownloadBook)
	r.GET("/hot", handler.GetHot)
	r.GET("/about", handler.About)

	if helper.SearchEnabled(){
		r.GET("/search", handler.SearchBook)
	}

	r.NoRoute(handler.NotFound)
	r.Static("/static", "static")
}

func InitHtml(r *gin.Engine) {
	r.SetFuncMap(template.FuncMap{"calHot": helper.CalHot})
	r.LoadHTMLGlob(filepath.Join(helper.GetRootPath(), "views/**/*"))
}
