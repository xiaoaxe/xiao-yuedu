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
	"github.com/githubao/xiao-yuedu/helper"
)

func TestSaveBook(t *testing.T) {
	InitDb()

	for i := range [50]int{} {
		book := Book{
			Name:     fmt.Sprintf("Effective Go %v", i+1),
			Content:  "最好的Go书籍之一",
			CreateAt: helper.JSONTime(time.Now()),
		}

		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		book.Create()

		log.Printf("process %v\n", i+1)
	}

}

func TestGetBooks(t *testing.T) {
	InitDb()

	books := FindAll()

	for _, book := range books {
		fmt.Println(book)
	}

}
