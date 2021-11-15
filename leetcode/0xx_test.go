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
var maxSubArrayGroup = []struct {
	inputs []int
	want   int
}{
	{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	{[]int{1}, 1},
	{[]int{0}, 0},
	{[]int{-1}, -1},
	{[]int{-1000}, -1000},
}

func TestMaxSubArray(t *testing.T) {
	for idx, test := range maxSubArrayGroup {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := maxSubArray(test.inputs)
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("失败！got:%#v want:%#v\n", got, test.want)
			}
		})
	}
}
