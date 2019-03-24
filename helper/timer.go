// 格式化时间格式
// author: baoqiang
// time: 2019/3/21 下午10:20
package helper

import (
	"time"
	"fmt"
)

// 血的教训，db里面的时间格式使用时间戳是最好的选择！
type JSONTime int64

//const LAYOUT = "2006-01-02 15:04:05.999"
const LAYOUT = "2006-01-02 15:04:05.999"

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Unix(0, int64(t)).Format(LAYOUT))
	return []byte(stamp), nil
}
