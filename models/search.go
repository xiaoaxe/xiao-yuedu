// 搜索结果model
// author: baoqiang
// time: 2019/3/22 下午3:17
package models

import "html/template"

type Search struct {
	Text  string
	Books []*Book
	PageTmpl template.HTML
}
