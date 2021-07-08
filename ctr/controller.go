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
		c.JSON(ExpressionIsEmpty, StatusText(ExpressionIsEmpty))
		return
	}

	if match, _ := regexp.MatchString("[0-9]+|[\\+\\*/\\-]+", expression); match != true {
		c.JSON(ThereAreOtherCharacters, StatusText(ThereAreOtherCharacters))
		return
	}

	expr, err := service.Calculate(expression)
	if err != nil {
		if err.Error() == "有连续操作符" {
			c.JSON(ContinuousOperator, StatusText(ContinuousOperator))
			return
		} else if err.Error() == "最后一位必须为操作数" {
			c.JSON(lastDigitIsNotNumber, StatusText(lastDigitIsNotNumber))
			return
		} else {
			c.JSON(InvalidDivisor, StatusText(InvalidDivisor))
			return
		}

	}
	c.JSON(Success, gin.H{Data: expr})
}

func Controller2(c *gin.Context) {
	exp := c.PostForm("exp")
	if exp == "" {
		c.JSON(ExpressionIsEmpty, StatusText(ExpressionIsEmpty))
		return
	}
	data, err := service.Caculator2(exp)

	if err != nil {
		c.JSON(FoundMistake, StatusText(FoundMistake))
		return
	}
	c.JSON(Success, StatusText1(Success, data))
}
