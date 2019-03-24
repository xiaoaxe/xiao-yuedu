// category 增删改查
// author: baoqiang
// time: 2019/3/22 下午5:42
package dal

import (
	"github.com/githubao/xiao-yuedu/models"
	"log"
)

func CreateCategory(c models.Category) error {
	return db.Model(&c).Create(&c).Error
}

func FindCategoryAll() []*models.Category {
	var categories []*models.Category

	err := db.Find(&categories).Error

	if err != nil {
		log.Printf("[ERROR] find all category failed: %+v", err)
		return nil
	}

	return categories
}
