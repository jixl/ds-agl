package leetcode

import (
	"reflect"
	"strconv"
	"testing"
)

var firstUniqCharCase = []struct {
	input string
	want  int
}{
	{"leetcode", 0},
	{"loveleetcode", 2},
	{"lovelove", -1},
}

func TestFirstUniqChar(t *testing.T) {
	for idx, test := range firstUniqCharCase {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := firstUniqChar(test.input)
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("失败！got:%#v want:%#v\n", got, test.want)
			}
			got = firstUniqChar1(test.input)
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("失败！got:%#v want:%#v\n", got, test.want)
			}
		})
	}
}

var canConstructCase = []struct {
	in1, in2 string
	want     bool
}{
	{"a", "b", false},
	{"aa", "ab", false},
	{"aa", "aab", true},
	{"baa", "aab", true},
}

func TestCanConstruct(t *testing.T) {
	for idx, test := range canConstructCase {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := canConstruct(test.in1, test.in2)
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("失败！got:%#v want:%#v\n", got, test.want)
			}
		})
	}
}
