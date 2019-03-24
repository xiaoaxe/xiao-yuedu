// 杂项函数
// author: baoqiang
// time: 2019/3/22 下午5:51
package helper

func SearchEnabled() bool {
	return GetConfig().Wukong.SearchEnable
}

func CalHot(downloadCount, viewCount int) int {
	return 3*downloadCount + 1*viewCount
}
