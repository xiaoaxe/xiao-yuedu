// 格式化时间格式
// author: baoqiang
// time: 2019/3/21 下午10:20
package helper

import (
	"time"
	"fmt"
)

type JSONTime time.Time

const LAYOUT = "2006-01-02 15:04:05.999999"

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(LAYOUT))
	return []byte(stamp), nil
}
