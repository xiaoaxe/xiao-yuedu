// 解析参数
// author: baoqiang
// time: 2019/3/21 下午8:16
package helper

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"log"
)

// param
func GetParamInt(c *gin.Context, key string) int {
	rs := c.Param(key)

	r, err := strconv.Atoi(rs)
	if err != nil {
		log.Printf("get int param failed: %v", err)
	}

	return r
}

func DefaultParamInt(c *gin.Context, key string, def int) int {
	rs := c.Param(key)

	r, err := strconv.Atoi(rs)
	if err != nil {
		log.Printf("get int param failed: %v", err)
		return def
	}

	return r
}

func GetParamInt64(c *gin.Context, key string) int64 {
	rs := c.Param(key)

	r, err := strconv.ParseInt(rs, 10, 64)
	if err != nil {
		log.Printf("get int param failed: %v", err)
	}

	return r
}

func DefaultParamInt64(c *gin.Context, key string, def int64) int64 {
	rs := c.Param(key)

	r, err := strconv.ParseInt(rs, 10, 64)
	if err != nil {
		log.Printf("get int param failed: %v", err)
		return def
	}

	return r
}

// query

func GetQueryInt(c *gin.Context, key string) int {
	rs := c.DefaultQuery(key, "")

	r, err := strconv.Atoi(rs)
	if err != nil {
		log.Printf("get int query failed: %v", err)
	}

	return r
}

func DefaultQueryInt(c *gin.Context, key string, def int) int {
	rs := c.Query(key)

	r, err := strconv.Atoi(rs)
	if err != nil {
		log.Printf("get int param failed: %v", err)
		return def
	}

	return r
}

func DefaultQueryInt64(c *gin.Context, key string, def int64) int64 {
	rs := c.Query(key)

	r, err := strconv.ParseInt(rs, 10, 64)
	if err != nil {
		log.Printf("get int param failed: %v", err)
		return def
	}

	return r
}
