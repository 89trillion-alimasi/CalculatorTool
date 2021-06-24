package service

import (
	"reflect"
	"testing"
)

func TestCalculate(t *testing.T) {
	type tests struct {
		exp  string
		want string
	}

	test := []tests{
		{exp: "1+2+3", want: "5"},
		{exp: "3+2-1+6/2", want: "7"},
		{exp: "3+2-1+6/0", want: ""},
	}
	for _, v := range test {
		got, _ := Calculate(v.exp)
		if !reflect.DeepEqual(got, v.want) {
			t.Error("expect:%v,got:%", v.want, got)
		}
	}

}
