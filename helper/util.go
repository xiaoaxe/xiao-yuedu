// 工具函数
// author: baoqiang
// time: 2019/3/21 下午8:31
package helper

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"os"
)

func GetRootPath() string {
	//path, _ := filepath.Abs(".")
	path, _ := os.Getwd()
	return path
}

func NotFound(v interface{}) gin.H {
	return gin.H{
		"code": 404,
		"msg":  fmt.Sprintf("%+v", v),
	}
}
