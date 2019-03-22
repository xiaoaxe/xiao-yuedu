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
	"github.com/githubao/xiao-yuedu/tools"
	"github.com/githubao/xiao-yuedu/models"
)

func TestSaveBook(t *testing.T) {
	InitDb()
	tools.InitSearch()

	for _,book := range books {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		Create(*book)

		log.Printf("process %v\n", book)
	}

}

func TestGetBooks(t *testing.T) {
	InitDb()

	books := FindBookAll()

	for _, book := range books {
		fmt.Println(book)
	}

}

var books = []*models.Book{
	{Name: "红楼梦", Content: "四大名著之一"},
	{Name: "Effective Go", Content: "很好的Golang学习书籍"},
	{Name: "The Golang Programming Language", Content: "Go编程语言"},
	{Name: "活着", Content: "学会忍受"},
	{Name: "明朝那些事儿", Content: "原来历史可以这么好玩儿"},
	{Name: "傲慢与偏见", Content: "关于女孩的独立和坚强"},
	{Name: "围城", Content: "每个人都是方鸿渐"},
}
