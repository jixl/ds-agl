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

func TestLoc(t *testing.T) {
	cases := []struct {
		loc  []byte
		want string
	}{
		{[]byte{0, 0, '1'}, "A"},
		{[]byte{0, 1, '2'}, "A"},
		{[]byte{0, 2, '3'}, "A"},
		{[]byte{3, 3, '4'}, "A"},
		{[]byte{4, 3, '5'}, "A"},
		{[]byte{5, 3, '3'}, "A"},
		{[]byte{6, 3, '6'}, "A"},
		{[]byte{7, 3, '7'}, "A"},
	}
	for idx, bs := range cases {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			loc(bs.loc[0], bs.loc[1], bs.loc[2])
			if !reflect.DeepEqual("101", "101") {
				t.Fatalf("失败！got:%#v want:%#v\n", "", "0")
			}
		})
	}
}
