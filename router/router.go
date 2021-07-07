package router

import (
	"CalculatorTool/ctr"
	"github.com/gin-gonic/gin"
)

// InitRouter 方法创建 gin 路由，并设置相关路由规则
func InitRouter() *gin.Engine {
	// 创建 gin router
	r := gin.Default()

	// 添加路由
	r.GET("/calculator", ctr.Controller)
	r.POST("/calculator2", ctr.Controller2)
	return r
}
