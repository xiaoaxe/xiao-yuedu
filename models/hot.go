// 热门的model
// author: baoqiang
// time: 2019/3/25 下午10:46
package models

import "html/template"

type Hot struct {
	Books    []*Book
	PageTmpl template.HTML
}