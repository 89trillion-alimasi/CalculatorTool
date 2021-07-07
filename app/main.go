package main

import (
	"CalculatorTool/env"
	"CalculatorTool/router"
	"fmt"
	"github.com/sirupsen/logrus"
)

// main 主执行方法
func main() {

	// 初始化 gin router
	r := router.InitRouter()
	port := env.Parse()
	// 启动 gin router
	err := r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		logrus.Error(err)
	}
}
