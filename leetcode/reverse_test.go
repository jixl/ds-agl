package leetcode

import (
	"reflect"
	"strconv"
	"testing"
)

var reverseStringCases = []struct{ in, want string }{
	{"Hello, world", "dlrow ,olleH"},
	{"Hello, 世界", "界世 ,olleH"},
	{"", ""},
}

func TestReverseRunes(t *testing.T) {
	for idx, test := range reverseStringCases {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := ReverseRunes(test.in)
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("失败！got:%#v want:%#v\n", got, test.want)
			}
		})
	}
}

var reverseWordsCases = []struct{ in, want string }{
	{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
	{" Hello,  世界 ", " ,olleH  界世 "},
	{"", ""},
}

func TestReverseWords(t *testing.T) {
	for idx, test := range reverseWordsCases {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := reverseWords(test.in)
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("失败！got:%#v want:%#v\n", got, test.want)
			}
		})
	}
}
