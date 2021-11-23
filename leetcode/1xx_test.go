package leetcode

import (
	"reflect"
	"strconv"
	"testing"
)

var connectCase = []struct{ in1, want []int }{
	{[]int{}, []int{}},
	{[]int{1, 2, 3, 4, 5, 6, 7}, []int{1, '#', 2, 3, '#', 4, 5, 6, 7, '#'}},
	{[]int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
		[]int{-1, '#', 0, 1, '#', 2, 3, 4, 5, '#', 6, 7, 8, 9, 10, 11, 12, 13, '#'}},
}

func Testconnect(t *testing.T) {
	for idx, test := range connectCase {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := connect(NewTree(test.in1))
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("失败！got:%#v want:%#v\n", got, test.want)
			}
		})
	}
}
