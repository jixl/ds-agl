package leetcode

import (
	"reflect"
	"strconv"
	"testing"
)

var checkInclusionCase = []struct {
	in1, in2 string
	want     bool
}{
	{"ab", "ab", true},
	{"ab", "eidbaooo", true},
	{"ab", "eidboaoo", false},
	{"hello", "ooolleoooleh", false},
	{"abc", "cccccbabbbaaaa", true},
	{"abc", "ccccbbbbaaaa", false},
	{"abc", "bbbca", true},
	// 3, 14
}

func TestCheckInclusion(t *testing.T) {
	for idx, test := range checkInclusionCase {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := checkInclusion(test.in1, test.in2)
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("失败！got:%#v want:%#v\n", got, test.want)
			}
		})
	}
}
