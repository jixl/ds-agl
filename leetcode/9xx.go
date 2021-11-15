package leetcode

// import "sort"

// 977. 有序数组的平方
// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
// 输入：nums = [-4,-1,0,3,10]
// 输出：[0,1,9,16,100]
// 解释：平方后，数组变为 [16,1,0,9,100]
// 排序后，数组变为 [0,1,9,16,100]
func sortedSquares(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	l, r := 0, len(nums)-1
	var rets = make([]int, len(nums))
	idx := len(nums) - 1
	for l <= r {
		if l == r {
			rets[idx] = nums[l] * nums[l]
			return rets
		}
		if nums[l]*nums[l] > nums[r]*nums[r] {
			rets[idx] = nums[l] * nums[l]
			l += 1
		} else {
			rets[idx] = nums[r] * nums[r]
			r -= 1
		}
		idx -= 1
	}
	return rets
}

// 环状替换 解法
func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	for s, i := 0, 0; i < n; s++ {
		idx, tmp := s, nums[s]
		for {
			i++
			idx = (idx + k) % n
			nums[idx], tmp = tmp, nums[idx]
			if s == idx {
				break
			}
		}
	}
}
