package leetcode

import "fmt"

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	for i, v := range numbers {
		f := target - v
		for l, r := i, n; l < r-1; {
			m := l + (r-l)>>1
			fmt.Printf("%d\t%d\t%d\t%d\t%d\n", i, l, r, f, m)
			if numbers[m] == f {
				return []int{i + 1, m + 1}
			} else if numbers[m] < f {
				l = m
			} else {
				r = m
			}
		}
	}
	return nil
}

// 217. 存在重复元素
func containsDuplicate(nums []int) bool {
	mp := make(map[int]bool, len(nums))
	for _, v := range nums {
		if mp[v] {
			return true
		}
		mp[v] = true
	}
	return false
}

func moveZeroes(nums []int) []int {
	n := len(nums)

	for z, d := 0, 0; z < n; z++ {
		if nums[z] != 0 {
			if z > d {
				nums[z], nums[d] = nums[d], nums[z]
			}
			d++
		}
	}
	return nums
}

func moveZeroes_sorted(nums []int) []int {
	n := len(nums)
	l, r := 0, n-1
	for l != r {
		if nums[r] == 0 {
			r--
		} else if nums[l] == 0 {
			nums[l], nums[r] = nums[r], nums[l]
			r--
		} else {
			l++
		}
		fmt.Printf("%d\t%d\t%d\n", l, r, nums)
	}

	if r == n-1 {
		return nums
	}

	for l = 0; l != r; {
		if nums[l] > nums[r] {
			nums[l], nums[r] = nums[r], nums[l]
			r--
		} else {
			l++
		}
		fmt.Printf("%d\t%d\t%d\n", l, r, nums)
	}
	return nums
}

// 566. 重塑矩阵
func matrixReshape(mat [][]int, r int, c int) [][]int {
	x, y := len(mat), len(mat[0])
	if x*y != r*c {
		return mat
	}

	var _x, _y int
	var remat = make([][]int, r)
	for i := 0; i < r; i++ {
		remat[i] = make([]int, c)
		for j := 0; j < c; j++ {
			if _y >= y {
				_x++
				_y = 0
			}
			remat[i][j] = mat[_x][_y]
			// remat[i] = append(remat[i], mat[_x][_y])
			_y++
		}
	}
	return remat
}

// 118 杨辉三角
func generate(numRows int) [][]int {
	tr := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		tr[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				tr[i][j] = 1
			} else {
				tr[i][j] = tr[i-1][j-1] + tr[i-1][j]
			}
		}
	}
	return tr
}
