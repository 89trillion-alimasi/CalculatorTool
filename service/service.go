package service

import (
	"CalculatorTool/model"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Calculate(s string) (string, error) {
	fmt.Println(s)
	s = strings.Replace(s, " ", "", -1)
	fmt.Println(s)
	stack := model.Stack{}
	preSign := '+'
	num := 0
	//if !lastisnum(s) {
	//	return -1
	//}
	for i, ch := range s {
		if isDigit(ch) {
			num = num*10 + int(ch-'0')
		}

		if !isDigit(ch) && i < len(s)-1 {
			if check(i+1, s) {
				return "", errors.New("")
			}
		}

		if !isDigit(ch) && ch != ' ' || i == len(s)-1 {
			if i == len(s)-1 {
				if !lastisnum(s) {
					return "", errors.New("最后一位必须为操作数")
				}
			}
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				if checkDividend(num) {
					return "", errors.New("被除数不能为零")
				}
				stack[len(stack)-1] /= num
			}
			preSign = ch
			num = 0
		}

	}
	var ans int
	for _, v := range stack {
		ans += v
	}
	return strconv.Itoa(ans), nil
}

//判断是否为数字
func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

//判断是不是有连续操作符
func check(i int, s string) bool {

	return s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/'
}

//判断被除数是否为0
func checkDividend(num int) bool {
	return num == 0
}

//判断最后一位是不是操作符
func lastisnum(s string) bool {
	i := len(s) - 1
	return '0' <= s[i] && s[i] <= '9'
}
