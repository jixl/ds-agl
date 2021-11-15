package leetcode

import "sort"

// 2032
func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	nums := make([]int, 0)

	m1 := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		m1[nums1[i]] = 1
	}

	for i := 0; i < len(nums2); i++ {
		num := m1[nums2[i]]
		if num != 0 && num != 2 {
			if num != -1 {
				nums = append(nums, nums2[i])
				m1[nums2[i]] = -1
			}
		} else {
			m1[nums2[i]] = 2
		}
	}

	for i := 0; i < len(nums3); i++ {
		num := m1[nums3[i]]
		if num != 0 && num != 3 {
			if num != -1 {
				nums = append(nums, nums3[i])
				m1[nums3[i]] = -1
			}
		} else {
			m1[nums3[i]] = 3
		}
	}
	return nums
}

// 2033
func minOperations(grid [][]int, x int) int {
	y := grid[0][0] % x
	m, n := len(grid), len(grid[0])
	nums := make([]int, 0, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j]%x != y {
				return -1
			}
			nums = append(nums, grid[i][j])
		}
	}

	sort.Ints(nums)
	mid := nums[len(nums)>>1]
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += abs(nums[i]-mid) / x
	}
	return sum
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// 链接：https://leetcode-cn.com/problems/stock-price-fluctuation
// 给你一支股票价格的数据流。数据流中每一条记录包含一个 时间戳 和该时间点股票对应的 价格 。

// 不巧的是，由于股票市场内在的波动性，股票价格记录可能不是按时间顺序到来的。某些情况下，有的记录可能是错的。如果两个有相同时间戳的记录出现在数据流中，前一条记录视为错误记录，后出现的记录 更正 前一条错误的记录。

// 请你设计一个算法，实现：

// 更新 股票在某一时间戳的股票价格，如果有之前同一时间戳的价格，这一操作将 更正 之前的错误价格。
// 找到当前记录里 最新股票价格 。最新股票价格 定义为时间戳最晚的股票价格。
// 找到当前记录里股票的 最高价格 。
// 找到当前记录里股票的 最低价格 。
// 请你实现 StockPrice 类：

// StockPrice() 初始化对象，当前无股票价格记录。
// void update(int timestamp, int price) 在时间点 timestamp 更新股票价格为 price 。
// int current() 返回股票 最新价格 。
// int maximum() 返回股票 最高价格 。
// int minimum() 返回股票 最低价格 。
//

// 示例 1：

// 输入：
// ["StockPrice", "update", "update", "current", "maximum", "update", "maximum", "update", "minimum"]
// [[], [1, 10], [2, 5], [], [], [1, 3], [], [4, 2], []]
// 输出：
// [null, null, null, 5, 10, null, 5, null, 2]

// 解释：
// StockPrice stockPrice = new StockPrice();
// stockPrice.update(1, 10); // 时间戳为 [1] ，对应的股票价格为 [10] 。
// stockPrice.update(2, 5);  // 时间戳为 [1,2] ，对应的股票价格为 [10,5] 。
// stockPrice.current();     // 返回 5 ，最新时间戳为 2 ，对应价格为 5 。
// stockPrice.maximum();     // 返回 10 ，最高价格的时间戳为 1 ，价格为 10 。
// stockPrice.update(1, 3);  // 之前时间戳为 1 的价格错误，价格更新为 3 。
//                           // 时间戳为 [1,2] ，对应股票价格为 [3,5] 。
// stockPrice.maximum();     // 返回 5 ，更正后最高价格为 5 。
// stockPrice.update(4, 2);  // 时间戳为 [1,2,4] ，对应价格为 [3,5,2] 。
// stockPrice.minimum();     // 返回 2 ，最低价格时间戳为 4 ，价格为 2 。
//

// 提示：
// 1 <= timestamp, price <= 109
// update，current，maximum 和 minimum 总 调用次数不超过 105 。
// current，maximum 和 minimum 被调用时，update 操作 至少 已经被调用过 一次 。

type StockPrice struct {
	prices  map[int]int
	current int
}

func Constructor() StockPrice {
	return StockPrice{
		prices:  make(map[int]int),
		current: 0,
	}
}

func (this *StockPrice) Update(timestamp int, price int) {
	this.prices[timestamp] = price
	if timestamp > this.current {
		this.current = timestamp
	}
}

func (this *StockPrice) Current() int {
	if this.current > 0 {
		return this.prices[this.current]
	}
	return 0
}

func (this *StockPrice) Maximum() int {
	max := 0
	for _, v := range this.prices {
		if max == 0 || max < v {
			max = v
		}
	}
	return max
}

func (this *StockPrice) Minimum() int {
	min := 0
	for _, v := range this.prices {
		if min == 0 || min > v {
			min = v
		}
	}
	return min
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
