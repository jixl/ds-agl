package leetcode

import (
	"reflect"
	"strconv"
	"testing"
)

// 53. 最大子序和
// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

var searchCase = []struct {
	inputs       []int
	target, want int
}{
	{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
	{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	{[]int{5}, 5, 0},
}

func TestSearch(t *testing.T) {
	for idx, test := range searchCase {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := search(test.inputs, test.target)
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("失败！got:%#v want:%#v\n", got, test.want)
			}
		})
	}
}
