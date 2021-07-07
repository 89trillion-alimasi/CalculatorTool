package model

type Stack []int

type Stack1 []string

//统计栈中元素的数量
func (s *Stack1) Len() int {
	return len(*s)
}

//将元素押入栈
func (s *Stack1) Push(value string) {
	*s = append(*s, value)
}

//获取栈顶元素
func (s Stack1) Top() string {
	return s[len(s)-1]
}

//弹出栈
func (s *Stack1) Pop() string {
	theStack := *s
	value := theStack[len(theStack)-1]
	*s = theStack[:len(theStack)-1]
	return value
}

func (s *Stack1) Isempty() bool {
	return len(*s) == 0
}
