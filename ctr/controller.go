package ctr

import (
	"CalculatorTool/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func Controller(c *gin.Context) {

	c.Request.URL.RawQuery = strings.ReplaceAll(c.Request.URL.RawQuery, "+", "%2b")
	//expression传进来的表达式
	expression := c.Query("expr")
	fmt.Println(expression)
	if expression == "" {
		c.String(400, "请输入表达式expr\n")
		return
	}

	expr, err := service.Calculate(expression)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.String(200, expr)
}
