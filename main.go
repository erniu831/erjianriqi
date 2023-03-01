package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个Gin实例
	router := gin.Default()

	// 定义路由处理函数
	router.GET("/", func(c *gin.Context) {
		remainingDays, remainingHours := timeRemainingUntilDate(time.Date(2023, time.June, 12, 9, 0, 0, 0, time.FixedZone("CST", 8*3600)))
		// res := fmt.Sprintf()
		c.String(200, "距离 2023年二级建造师 考试还有%s, 换算成小时还有 %s", remainingDays, remainingHours)
	})

	// 启动服务器并监听端口
	router.Run(":8080")
}

func timeRemainingUntilDate(date time.Time) (string, string) {
	now := time.Now().UTC()
	remainingDuration := date.Sub(now)
	remainingDays := int(remainingDuration.Hours() / 24)
	remainingHours := int(remainingDuration.Hours()) % 24
	return fmt.Sprintf("%d 天 %d 小时", remainingDays, remainingHours), fmt.Sprintf("%d 小时", int(remainingDuration.Hours()))
}
