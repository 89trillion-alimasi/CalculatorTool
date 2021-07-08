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

	expr, err, t := service.Calculate(expression)
	if err != nil {
		switch t {
		case 1:
			c.JSON(ContinuousOperator, StatusText(ContinuousOperator))
			return
		case 2:
			c.JSON(lastDigitIsNotNumber, StatusText(lastDigitIsNotNumber))
			return
		case 3:
			c.JSON(InvalidDivisor, StatusText(InvalidDivisor))
			return
		}
	}
	c.JSON(Success, StatusText1(Success, expr))
}

func Controller2(c *gin.Context) {
	exp := c.PostForm("exp")
	if exp == "" {
		c.JSON(ExpressionIsEmpty, StatusText(ExpressionIsEmpty))
		return
	}
	data, err, t := service.Caculator2(exp)

	if err != nil {
		switch t {
		case 1:
			c.JSON(ContinuousOperator, StatusText(ContinuousOperator))
			return
		case 2:
			c.JSON(lastDigitIsNotNumber, StatusText(lastDigitIsNotNumber))
			return
		case 3:
			c.JSON(InvalidDivisor, StatusText(InvalidDivisor))
			return
		}
	}

	c.JSON(Success, StatusText1(Success, data))
}
