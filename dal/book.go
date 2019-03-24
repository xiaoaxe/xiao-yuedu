// book model
// author: baoqiang
// time: 2019/3/21 下午6:36
package dal

import (
	"log"
	"github.com/githubao/xiao-yuedu/models"
	"github.com/githubao/xiao-yuedu/helper"
	"time"
	"fmt"
)

func CreateBook(b models.Book) error {
	b.CreatedAt = helper.JSONTime(time.Now().UnixNano())
	b.UpdatedAt = b.CreatedAt
	return db.Model(&b).Create(&b).Error
}

// UPDATE start
func UpdateBookField(b models.Book, attr string) error {
	addUpdate(&b)
	return db.Model(&b).Select(addUpdateField(attr)).UpdateColumns(b).Error
}

func AddView(b models.Book) error {
	addUpdate(&b)
	b.ViewCount += 1
	return db.Model(&b).Select(addUpdateField("view_count")).UpdateColumns(b).Error
}

func AddDownload(b models.Book) error {
	addUpdate(&b)
	b.DownloadCount += 1
	return db.Model(&b).Select(addUpdateField("download_count")).UpdateColumns(b).Error
}

func addUpdate(b *models.Book) {
	b.UpdatedAt = helper.JSONTime(time.Now().UnixNano())
}

func addUpdateField(attr string) []string {
	return []string{attr, "updated_at"}
}

// UPDATE end

func CountBook() int64 {
	var count int64
	err := db.Model(&models.Book{}).Count(&count).Error
	if err != nil {
		log.Printf("[ERROR] count failed: %v", err)
		return 0
	}
	return count
}

// FIND start
func FindBookOne(bid int64) *models.Book {
	var book models.Book

	err := db.Where("id=?", bid).Find(&book).Error

	if err != nil {
		log.Printf("[ERROR] find bid failed: %d, %+v", bid, err)
		return nil
	}

	return &book
}

func FindBookIds(bids []int64) []*models.Book {
	var books []*models.Book

	err := db.Where("id in (?)", bids).Find(&books).Error

	if err != nil {
		log.Printf("[ERROR] find bids failed: %d, %+v", bids, err)
		return nil
	}

	return books
}

func FindBookAll() []*models.Book {
	var books []*models.Book

	err := db.Find(&books).Error

	if err != nil {
		log.Printf("[ERROR] find all book failed: %+v", err)
		return nil
	}

	return books
}

func FindBookOrder(pageNum int, perPage int) []*models.Book {
	var books []*models.Book

	offset := (pageNum - 1) * perPage

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

func FindHot(pageNum int, perPage int) []*models.Book {
	var books []*models.Book

	offset := (pageNum - 1) * perPage

	orderValueFmt := "(%d * download_count + %d * view_count) DESC"
	orderValue := fmt.Sprintf(orderValueFmt, 3, 1)

	err := db.Model(&books).
		Order(orderValue).
		Limit(perPage).
		Offset(offset).
		Find(&books).Error

	if err != nil {
		log.Printf("[ERROR] find all failed: %+v", err)
		return nil
	}

	return books
}

// FIND end

// deprecated update all
//func Update(b models.Book) error {
//	return db.Model(&b).Save(b).Error
//}
