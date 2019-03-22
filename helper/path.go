// 路径名相关的工具函数
// author: baoqiang
// time: 2019/3/21 下午8:31
package helper

import (
	"os"
	"path/filepath"
)

func GetRootPath() string {
	//path, _ := filepath.Abs(".")
	path, _ := os.Getwd()
	return path
}

func GetDictionaryFile() string {
	return filepath.Join(GetRootPath(), config.Wukong.DictFile)
}
