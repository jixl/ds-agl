package leetcode

import (
	"reflect"
	"strconv"
	"testing"
)

var lengthOfLongestSubstringCases = []struct {
	inputs string
	want   int
}{
	{"", 0}, {" ", 1}, {"ab", 2}, {"aaaa", 1},
	{"abcabc", 3}, {"pwwkew", 3},
	{"tmmzuxt", 5},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for idx, test := range lengthOfLongestSubstringCases {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := lengthOfLongestSubstring(test.inputs)
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("失败！got:%#v want:%#v\n", got, test.want)
			}
		})
	}
}

var mergeTwoListsCases = []struct{ in1, in2, want []int }{
	{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
}

func TestMergeTwoLists(t *testing.T) {
	for idx, test := range mergeTwoListsCases {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := mergeTwoLists(NewList(test.in1), NewList(test.in2))
			if !reflect.DeepEqual(got.toArray(), test.want) {
				t.Fatalf("失败！got:%#v want:%#v\n", got.toArray(), test.want)
			}
		})
	}
}

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
