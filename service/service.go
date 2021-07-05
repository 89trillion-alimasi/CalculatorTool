package service

import (
	"CalculatorTool/model"
	"errors"
	"strconv"
)

func Calculate(s string) (int, error) {

	stack := model.Stack{}
	preSign := '+'
	num := ""

	//判断表达式是否正确
	for i, ch := range s {
		//是否是连续操作符
		if !isDigit(ch) && i < len(s)-1 {
			if check(i+1, s) {
				return -1, errors.New("有连续操作符")
			}
		} else if i == len(s)-1 { //最后一位必须是操作数
			if !lastisnum(s) {
				return -1, errors.New("最后一位必须为操作数")
			}
		}
	}

	for i, ch := range s {
		if isDigit(ch) {
			num = num + string(ch)
		}
		if !isDigit(ch) && ch != ' ' || i == len(s)-1 {
			nums, _ := strconv.Atoi(num)
			switch preSign {
			case '+':
				stack = append(stack, nums)
			case '-':
				stack = append(stack, -nums)
			case '*':
				stack[len(stack)-1] *= nums
			case '/':
				if checkDividend(nums) {
					return -1, errors.New("被除数不能为零")
				}
				stack[len(stack)-1] /= nums
			}
			preSign = ch
			nums = 0
			num = ""
		}

	}
	var ans int
	for _, v := range stack {
		ans += v
	}
	return ans, nil
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
