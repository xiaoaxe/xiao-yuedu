// 下载文件
// author: baoqiang
// time: 2019/3/22 下午5:29
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/githubao/xiao-yuedu/dal"
	"fmt"
	"net/http"
	"github.com/githubao/xiao-yuedu/consts"
	"github.com/githubao/xiao-yuedu/helper"
	"github.com/githubao/xiao-yuedu/models"
	"net/url"
	"path"
	"os"
	"log"
	"errors"
)

func DownloadBook(c *gin.Context) {
	bid := helper.GetQueryInt64(c, "book_id")

	book := dal.FindBookOne(bid)

	if book == nil {
		msg := fmt.Sprintf("book NOT FOUND: %v", bid)
		c.HTML(http.StatusNotFound, consts.ErrorTmpl, msg)
		return
	}

	// get file path
	filename,err := buildFilename(book)
	if err != nil{
		c.HTML(http.StatusNotFound,consts.ErrorTmpl,err.Error())
		return
	}

	// set header
	setDownloadHeader(c, filename)

	// incr download count
	dal.AddDownload(*book)

	// download
	c.File(filename)
}

func buildFilename(book *models.Book) (string,error) {
	rootFilePath := helper.GetConfig().Yuedu.DlPath
	major, sub := book.Category[0:3], book.Category[3:7]
	filename := fmt.Sprintf("%s/%s/%s/%s.%s", rootFilePath, major, sub,
		book.Name, book.Format)

	file,err := os.Open(filename)
	defer file.Close()

	if err != nil{
		log.Printf("[ERROR] file not found: %s", filename)
		return "",errors.New("file not found")
	}


	return filename,nil
}

func setDownloadHeader(c *gin.Context, filename string) {
	// 设置浏览器是否为直接下载文件，且为浏览器指定下载文件的名字
	c.Header("Content-Disposition", "attachment; filename="+url.QueryEscape(path.Base(filename)))
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Expires", "0")
	// 如果缓存过期了，会再次和原来的服务器确定是否为最新数据，而不是和中间的proxy
	c.Header("Cache-Control", "must-revalidate")
	c.Header("Pragma", "public")
}
