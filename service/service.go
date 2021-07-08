package service

import (
	"CalculatorTool/model"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Calculate(s string) (int, error, interface{}) {
	fmt.Printf(s)
	s = strings.Replace(s, " ", "", -1)
	stack := model.Stack{}
	preSign := '+'
	num := ""

	//判断表达式是否正确
	for i, ch := range s {
		//是否是连续操作符
		if !isDigit(ch) && i < len(s)-1 {
			if check(i+1, s) {
				return -1, errors.New("有连续操作符"), 1
			}
		} else if i == len(s)-1 { //最后一位必须是操作数
			if !Islastnum(s) {
				return -1, errors.New("最后一位必须为操作数"), 2
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
					return -1, errors.New("被除数不能为零"), 3
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
	return ans, nil, nil
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
func Islastnum(s string) bool {
	i := len(s) - 1
	return '0' <= s[i] && s[i] <= '9'
}

//后缀表达式
func Caculator2(s string) (data int, err error, t interface{}) {
	if err, t := Judgement(s); !err {
		if t == 2 {
			return -1, errors.New("最后一位是操作符"), 2
		}
		return -1, errors.New("有连续操作符"), 1
	}
	suffix := MixTosuffix(s)

	var curStack model.Stack1
	for _, v := range suffix {
		curChar := v
		if i, ero := strconv.Atoi(v); ero == nil {
			curStack.Push(strconv.Itoa(i))
		} else {
			v1 := curStack.Pop()
			v2 := curStack.Pop()
			num1, _ := strconv.Atoi(v1)
			num2, _ := strconv.Atoi(v2)
			switch curChar {
			case "+":
				curStack.Push(strconv.Itoa(num2 + num1))
			case "-":
				curStack.Push(strconv.Itoa(num2 - num1))
			case "*":
				curStack.Push(strconv.Itoa(num2 * num1))
			case "/":
				if num1 == 0 {
					return -1, errors.New("除数不能为零"), 3
				}
				curStack.Push(strconv.Itoa(num2 / num1))
			}
		}
	}
	ret := curStack.Top()
	result, _ := strconv.Atoi(ret)
	return result, nil, nil
}

//中缀表达式转后缀表达式
func MixTosuffix(exp string) []string {
	var curStack model.Stack1
	var suffix []string
	expLen := len(exp)
	for i := 0; i < expLen; i++ {
		char := string(exp[i])
		if char == " " {
			continue
		} else if unicode.IsDigit(rune(exp[i])) {
			j := i
			digit := ""
			for ; j < expLen && unicode.IsDigit(rune(exp[j])); j++ {
				digit += string(exp[j])
			}
			suffix = append(suffix, digit)
			i = j - 1
		} else {
			for !curStack.Isempty() {
				top := curStack.Top()
				if priority(top, char) {
					break
				}
				suffix = append(suffix, top)
				curStack.Pop()
			}
			curStack.Push(char)
		}
	}
	for !curStack.Isempty() {
		data := curStack.Pop()
		suffix = append(suffix, data)
	}
	return suffix
}

//判断优先级
func priority(top, char string) bool {
	switch top {
	case "+", "-":
		if char == "*" || char == "/" {
			return true
		}
	}
	return false
}

//判断表达式是否正确
func Judgement(exp string) (bool, interface{}) {
	//如果算数表达式的第一个位置的元素不是数字或者最后一个位置的元素不是数字返回false
	if !unicode.IsDigit(rune(exp[0])) || !unicode.IsDigit(rune(exp[len(exp)-1])) {
		return false, 1
	}
	flag := false
	for _, char := range exp {
		//如果字符为空，跳过判断
		if char == ' ' {
			continue
		} else if !unicode.IsDigit(rune(char)) && flag == false {
			//如果当前不是数字那么检查下一个是否是操作符
			flag = true
		} else if !unicode.IsDigit(rune(char)) && flag == true {
			//不能是连续操作符
			return false, 2
		} else if unicode.IsDigit(rune(char)) && flag == true {
			//如果当前是数字，看下面是否还有数字和判断下个是否是操作符
			flag = false
		}
	}
	return true, nil
}
