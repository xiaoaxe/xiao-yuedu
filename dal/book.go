// book model
// author: baoqiang
// time: 2019/3/21 下午6:36
package dal

import (
	"log"
	"github.com/githubao/xiao-yuedu/models"
)

func Create(b models.Book) error {
	return db.Model(&b).Create(b).Error
	return nil
}

func UpdateBookField(b models.Book, attr string) error {
	return db.Model(&b).Select(attr).Updates(b).Error
}

func CountBook() int64 {
	var count int64
	err := db.Model(&models.Book{}).Count(&count).Error
	if err != nil {
		log.Printf("[ERROR] count failed: %v", err)
		return 0
	}
	return count
}

// FIND START
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
	var book []*models.Book

	err := db.Find(&book).Error

	if err != nil {
		log.Printf("[ERROR] find all failed: %+v", err)
		return nil
	}

	return book
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

// FIND END

// deprecated update all
func Update(b models.Book) error {
	return db.Model(&b).Save(b).Error
}
