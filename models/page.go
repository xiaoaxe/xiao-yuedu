// 分页返回的数据
// author: baoqiang
// time: 2019/3/25 下午10:01
package models

import "html/template"

type Page struct {
	Books    []*Book
	PageTmpl template.HTML
}
