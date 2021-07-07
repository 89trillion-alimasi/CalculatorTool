package ctr

import (
	"CalculatorTool/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
)

func Controller(c *gin.Context) {

	c.Request.URL.RawQuery = strings.ReplaceAll(c.Request.URL.RawQuery, "+", "%2b")
	//expression传进来的表达式
	expression := c.Query("expr")
	fmt.Println(expression)

	if expression == "" {
		c.JSON(ExpressionIsEmpty, gin.H{Message: "请输入表达式expr"})
		return
	}

	if match, _ := regexp.MatchString("[0-9]+|[\\+\\*/\\-]+", expression); match != true {
		c.JSON(ThereAreOtherCharacters, gin.H{Message: "有除了数字和字符串以外的字符"})
		return
	}

	expr, err := service.Calculate(expression)
	if err != nil {
		if err.Error() == "有连续操作符" {
			c.JSON(ContinuousOperator, gin.H{Message: err.Error()})
			return
		} else if err.Error() == "最后一位必须为操作数" {
			c.JSON(lastDigitIsNotNumber, gin.H{Message: err.Error()})
			return
		} else {
			c.JSON(InvalidDivisor, gin.H{Message: err.Error()})
			return
		}

	}
	c.JSON(Success, gin.H{Message: expr})
}

func Controller2(c *gin.Context) {
	exp := c.PostForm("exp")
	if exp == "" {
		c.JSON(ExpressionIsEmpty, gin.H{Message: StatusText(ExpressionIsEmpty)})
		return
	}
	data, err := service.Caculator2(exp)

	if err != nil {
		c.JSON(FoundMistake, gin.H{Message: StatusText(FoundMistake)})
		return
	}
	c.JSON(Success, gin.H{Message: StatusText(Success),
		Data: data})
}
