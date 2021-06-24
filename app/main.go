package main

import (
	"CalculatorTool/router"
	"github.com/sirupsen/logrus"
)

// main 主执行方法
func main() {

	// 初始化 gin router
	r := router.InitRouter()

	// 启动 gin router
	err := r.Run()
	if err != nil {
		logrus.Error(err)
	}
}
