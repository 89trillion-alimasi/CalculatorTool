package ctr

import (
	"CalculatorTool/service"
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
)

func Controller(c *gin.Context) {

	c.Request.URL.RawQuery = strings.ReplaceAll(c.Request.URL.RawQuery, "+", "%2b")
	//expression传进来的表达式
	expression := c.Query("expr")
	//fmt.Println(expression)

	if expression == "" {
		c.JSON(400, gin.H{"Message": "请输入表达式expr"})
		return
	}

	if match, _ := regexp.MatchString("[0-9]+|[\\+\\*/\\-]+", expression); match != true {
		c.JSON(401, gin.H{"message": "有除了数字和字符串以外的字符"})
		return
	}

	expr, err := service.Calculate(expression)
	if err != nil {
		if err.Error() == "有连续操作符" {
			c.JSON(402, gin.H{"Message": err.Error()})
			return
		} else if err.Error() == "最后一位必须为操作数" {
			c.JSON(403, gin.H{"Message": err.Error()})
			return
		} else {
			c.JSON(404, gin.H{"Message": err.Error()})
			return
		}

	}
	c.JSON(200, gin.H{"data": expr})
}
