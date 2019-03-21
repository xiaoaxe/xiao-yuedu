// book model
// author: baoqiang
// time: 2019/3/21 下午6:36
package dal

import (
	"log"
	"encoding/json"
	"github.com/githubao/xiao-yuedu/helper"
)

type Book struct {
	Id         int64           `gorm:"id,primary_key"`
	Name       string          `gorm:"name"`
	Format     string          `gorm:"format"`
	Content    string          `gorm:"content"`
	Tags       string          `gorm:"tags"`
	Categories string          `gorm:"categories"`
	ViewCount  int             `gorm:"view_count"`
	DlCount    int             `gorm:"dl_count"`
	PanUrl     string          `gorm:"pan_url"`
	CreateAt   helper.JSONTime `gorm:"created_at"`
	UpdateAt   helper.JSONTime `gorm:"updated_at"`
}

func (b *Book) TableName() string {
	return "book"
}

func (b *Book) Create() error {
	db := GetDb()
	return db.Create(b).Error
}

func (b *Book) Update() error {
	db := GetDb()
	return db.Save(b).Error
}

func (b *Book) String() string {
	jdata, _ := json.Marshal(b)
	return string(jdata)
}

func FindOne(bid int64) *Book {
	db := GetDb()

	var book Book

	err := db.Where("id=?", bid).Find(&book).Error

	if err != nil {
		log.Printf("[ERROR] find bid failed: %d", bid)
		return nil
	}

	return &book
}

func FindAll() []*Book {
	db := GetDb()

	var book []*Book

	err := db.Find(&book).Error

	if err != nil {
		log.Printf("[ERROR] find all failed: %v", err)
		return nil
	}

	return book
}

func FindOrder(page int64, perPage int64) []*Book {
	db := GetDb()

	var books []*Book

	offset := (page - 1) * perPage

	err := db.Model(&books).
		Order("created_at DESC").
		Limit(perPage).
		Offset(offset).
		Find(&books).Error

	if err != nil {
		log.Printf("[ERROR] find all failed: %+v", err)
		return nil
	}

	return books

}
