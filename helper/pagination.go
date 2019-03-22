// 分页
// author: baoqiang
// time: 2019/3/22 下午3:42
package helper

import "fmt"

func AssemblePageMsg(pageNum int, total, perPage int) string {
	maxPage := calcMaxPage(total, perPage)

	msg := fmt.Sprintf("page too large: %d, limit is 1", pageNum)
	if maxPage > 1 {
		msg = fmt.Sprintf("%s-%d", msg, maxPage)
	}

	return msg
}

func calcMaxPage(total int, perPage int) int {
	cnt, mod := total/perPage, total%perPage
	if mod == 0 {
		return cnt
	} else {
		return cnt + 1
	}
}
