// 书籍搜索
// author: baoqiang
// time: 2019/3/22 下午1:51
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/githubao/xiao-yuedu/tools"
	"github.com/githubao/xiao-yuedu/helper"
	"net/http"
	"github.com/githubao/xiao-yuedu/consts"
	"fmt"
	"github.com/githubao/xiao-yuedu/dal"
	"github.com/githubao/xiao-yuedu/models"
	"github.com/huichen/wukong/types"
)

func SearchBook(c *gin.Context) {
	text := c.DefaultQuery("text", "")
	pageNum := helper.DefaultQueryInt(c, "p", 1)
	//io.Copy(ioutil.Discard, strings.NewReader(fmt.Sprintf("%d", pid)))

	sret := tools.SearchDocument(text)

	if sret.NumDocs <= 0 {
		msg := fmt.Sprintf("search not found: %s", text)
		c.HTML(http.StatusOK, consts.ErrorTmpl, msg)
		return
	}

	// 处理数据，如果页数太大数据为空返回错误页
	var bids = filterPage(sret, pageNum)
	if len(bids) == 0 {
		msg := helper.AssemblePageMsg(pageNum, len(sret.Docs), consts.PerPage)
		c.HTML(http.StatusOK, consts.ErrorTmpl, msg)
		return
	}

	books := dal.FindBookIds(bids)

	c.HTML(http.StatusOK, consts.SearchTmpl, models.Search{
		Text:  text,
		Books: books,
	})
}

func filterPage(sret types.SearchResponse, pageNum int) []int64 {
	var bids []int64

	var filtered []types.ScoredDocument

	start := (pageNum - 1) * consts.PerPage
	end := pageNum * consts.PerPage
	counts := sret.NumDocs

	// 不同的起始位置
	if counts >= end {
		filtered = sret.Docs[start:end]
	} else if counts >= start {
		filtered = sret.Docs[start:counts]
	} else {
		return bids
	}

	for _, doc := range filtered {
		bids = append(bids, int64(doc.DocId))
	}

	return bids
}
