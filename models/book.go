// 书籍
// author: baoqiang
// time: 2019/3/22 下午2:48
package models

import (
	"github.com/githubao/xiao-yuedu/helper"
	"encoding/json"
)

type Book struct {
	Id        int64           `gorm:"id,primary_key"`
	Name      string          `gorm:"name"`
	Format    string          `gorm:"format"`
	Content   string          `gorm:"content"`
	Tags      string          `gorm:"tags"`
	Category  string          `gorm:"category"`
	ViewCount int             `gorm:"view_count"`
	DlCount   int             `gorm:"dl_count"`
	PanUrl    string          `gorm:"pan_url"`
	CreateAt  helper.JSONTime `gorm:"created_at"`
	UpdateAt  helper.JSONTime `gorm:"updated_at"`
}

func (b *Book) String() string {
	jdata, _ := json.Marshal(b)
	return string(jdata)
}
