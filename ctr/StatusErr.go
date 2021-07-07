package ctr

const (
	Message                 = "Message"
	Data                    = "Data"
	ExpressionIsEmpty       = 400
	ThereAreOtherCharacters = 401
	ContinuousOperator      = 402
	lastDigitIsNotNumber    = 403
	InvalidDivisor          = 404
	FoundMistake            = 405
	Success                 = 200
)

var statusText = map[int]string{
	ExpressionIsEmpty:       "表达式为空",
	ThereAreOtherCharacters: "有其他的字符",
	ContinuousOperator:      "有连续操作符",
	lastDigitIsNotNumber:    "最后一位不是操作符",
	InvalidDivisor:          "除数不能为0",
	Success:                 "成功",
	FoundMistake:            "计算出现错误",
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}
