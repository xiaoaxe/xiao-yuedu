// 搜索测试
// author: baoqiang
// time: 2019/3/22 下午1:37
package tools

import (
	"testing"
	"fmt"
	"github.com/githubao/xiao-yuedu/dal"
)

func TestSearch(t *testing.T) {
	dal.InitDb()
	InitSearch()

	for _, text := range []string{
		"明朝那些事儿", "历史原来", "历史", "忍受",
		"每个", "围城", "每个人",
	} {
		resp := SearchDocument(text)
		fmt.Println(resp)
	}

}
