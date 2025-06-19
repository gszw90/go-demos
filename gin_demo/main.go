package main

import (
	"github.com/gin-gonic/gin"

	"demos/gin_demo/middleware"
)

func main() {
	r := gin.Default()
	r.Use(
		middleware.AccessLog(),
		middleware.Params,
	)
	r.POST("/raw", RawHandler)

	r.Run(":8080")
}

// RawHandler 处理原始请求
func RawHandler(c *gin.Context) {

	// 获取原始请求体
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to get raw data"})
		return
	}
	header := c.Request.Header
	headerSignature := c.GetHeader(middleware.HITPAY_CALLBACK_HEADER_SIGNATURE_KEY)

	// 处理请求体
	bodyParam, _ := c.Get(middleware.MiddlewareParamsBodyCtxKey)
	// 返回响应
	c.JSON(200, gin.H{
		"message":   "Raw request handled successfully",
		"data":      string(body),
		"signature": headerSignature,
		"header":    header,
		"body":      bodyParam,
	})
}
