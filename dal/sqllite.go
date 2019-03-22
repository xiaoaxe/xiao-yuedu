// 使用文件数据库
// author: baoqiang
// time: 2019/3/21 下午6:39
package dal

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/githubao/xiao-yuedu/helper"
)

var db *gorm.DB

func InitDb() *gorm.DB {
	client, err := gorm.Open("sqlite3", helper.GetConfig().Yuedu.DbFile)
	if err != nil {
		panic("init db failed")
	}
	db = client
	return db
}
