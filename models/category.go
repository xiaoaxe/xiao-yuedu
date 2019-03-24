// 类别
// author: baoqiang
// time: 2019/3/22 下午5:40
package models

import (
	"github.com/githubao/xiao-yuedu/helper"
	"encoding/json"
)

type Category struct {
	Id        int64           `gorm:"id,primary_key"`
	Category  string          `gorm:"category"`
	Major     string          `gorm:"major"`
	Sub       string          `gorm:"sub"`
	CreatedAt helper.JSONTime `gorm:"created_at"`
	UpdatedAt helper.JSONTime `gorm:"updated_at"`
}

func (c *Category) String() string {
	jdata, _ := json.Marshal(c)
	return string(jdata)
}
