// 插入一条数据
// author: baoqiang
// time: 2019/3/21 下午8:46
package dal

import (
	"testing"
	"fmt"
	"time"
	"math/rand"
	"log"
	"github.com/githubao/xiao-yuedu/models"
)

func aTestSaveBook(t *testing.T) {
	InitDb()

	for _, book := range books {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		CreateBook(*book)

		log.Printf("process %v\n", book)
	}
}

func aTestSaveCategory(t *testing.T) {
	InitDb()

	for _, category := range categories {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		err := CreateCategory(*category)
		if err != nil {
			panic(err)
		}

		log.Printf("process %v\n", category)
	}
}

func aTestGetBooks(t *testing.T) {
	InitDb()

	books := FindBookAll()

	for _, book := range books {
		fmt.Println(book)
	}
}

func aTestGetCategory(t *testing.T) {
	InitDb()

	categories := FindCategoryAll()

	for _, cate := range categories {
		fmt.Println(cate)
	}
}

func TestUpdateBook(t *testing.T) {
	InitDb()

	for _, book := range uptBooks {
		err := UpdateBookField(book,"category")
		err = UpdateBookField(book,"format")
		if err != nil {
			panic(err)
		}

		log.Printf("process %v\n", book)
	}
}

var books = []*models.Book{
	{Name: "红楼梦", Content: "四大名著之一"},
	{Name: "Effective_Go", Content: "很好的Golang学习书籍"},
	{Name: "The_Golang_Programming_Language", Content: "Go编程语言"},
	{Name: "活着", Content: "学会忍受"},
	{Name: "明朝那些事儿", Content: "原来历史可以这么好玩儿"},
	{Name: "傲慢与偏见", Content: "关于女孩的独立和坚强"},
	{Name: "围城", Content: "每个人都是方鸿渐"},
}

var categories = []*models.Category{
	{Category: "1000001", Major: "其他", Sub: "其他"},
	{Category: "1010001", Major: "编程", Sub: "Golang"},
	{Category: "1020001", Major: "文学", Sub: "中国古典"},
	{Category: "1020002", Major: "文学", Sub: "外国名著"},
}

var uptBooks = []models.Book{
	{Id: 1, Category: "1020001",Format:"pdf"},
	{Id: 2, Category: "1010001",Format:"pdf"},
	{Id: 3, Category: "1010001",Format:"pdf"},
	{Id: 4, Category: "1020001",Format:"pdf"},
	{Id: 5, Category: "1020001",Format:"pdf"},
	{Id: 6, Category: "1020002",Format:"pdf"},
	{Id: 7, Category: "1020001",Format:"pdf"},
}
