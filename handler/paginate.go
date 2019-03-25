// 分页实现，cp from https://studygolang.com/articles/11649
// author: baoqiang
// time: 2019/3/25 下午10:21
package handler

import (
	"math"
	"fmt"
	"strings"
	"html/template"
	"github.com/githubao/xiao-yuedu/helper"
	"github.com/githubao/xiao-yuedu/consts"
)

//Pagination 分页器
type Pagination struct {
	CurrentPage int
	TotalPage   int
	Url         string
}

func BuildPaginateHtml(url string, currentPage, count, perPage int) template.HTML {
	totalPage := helper.CalcMaxPage(int(count), perPage)

	tempStr := NewPagination(url, currentPage, totalPage).Pages()

	return template.HTML(tempStr)
}

//NewPagination 新建分页器
func NewPagination(url string, currentPage int, totalPage int) *Pagination {
	return &Pagination{
		CurrentPage: currentPage,
		TotalPage:   totalPage,
		Url:         url,
	}
}

//Pages 渲染生成html分页标签
func (p *Pagination) Pages() string {
	//计算总页数
	currentPage := p.CurrentPage
	totalPage := p.TotalPage

	//首页链接
	var firstLink string
	//上一页链接
	var prevLink string
	//下一页链接
	var nextLink string
	//末页链接
	var lastLink string
	//中间页码链接
	var pageLinks []string

	//首页和上一页链接
	if currentPage > 1 {
		firstLink = fmt.Sprintf(`<li><a href="%s">首页</a></li>`, p.pageURL(1))
		prevLink = fmt.Sprintf(`<li><a href="%s">上一页</a></li>`, p.pageURL(currentPage-1))
	} else {
		firstLink = `<li class="disabled"><a href="#">首页</a></li>`
		prevLink = `<li class="disabled"><a href="#">上一页</a></li>`
	}

	//末页和下一页
	if currentPage < totalPage {
		lastLink = fmt.Sprintf(`<li><a href="%s">末页</a></li>`, p.pageURL(totalPage))
		nextLink = fmt.Sprintf(`<li><a href="%s">下一页</a></li>`, p.pageURL(currentPage+1))
	} else {
		lastLink = `<li class="disabled"><a href="#">末页</a></li>`
		nextLink = `<li class="disabled"><a href="#">下一页</a></li>`
	}

	//生成中间页码链接
	pageLinks = make([]string, 0, 10)
	startPos := currentPage - consts.PageDelta
	endPos := currentPage + consts.PageDelta
	if startPos < 1 {
		endPos = endPos + int(math.Abs(float64(startPos))) + 1
		startPos = 1
	}
	if endPos > totalPage {
		endPos = totalPage
	}
	for i := startPos; i <= endPos; i++ {
		var s string
		if i == currentPage {
			s = fmt.Sprintf(`<li class="active"><a href="%s">%d</a></li>`, p.pageURL(i), i)
		} else {
			s = fmt.Sprintf(`<li><a href="%s">%d</a></li>`, p.pageURL(i), i)
		}
		pageLinks = append(pageLinks, s)
	}

	return fmt.Sprintf(`<nav style="text-align: center" class="footer"><ul class="pagination" >%s%s%s%s%s</ul></nav>`,
		firstLink, prevLink, strings.Join(pageLinks, ""), nextLink, lastLink)
}

//pageURL 生成分页url
func (p *Pagination) pageURL(newerPage int) string {
	currentPage := p.CurrentPage

	if strings.Contains(p.Url, "pages") {
		return fmt.Sprintf("/pages/%d", newerPage)
	} else {
		url := helper.AddPagePath(p.Url)
		current := fmt.Sprintf("p=%d", currentPage)
		newer := fmt.Sprintf("p=%d", newerPage)
		return strings.Replace(url, current, newer, 1)
	}
}
