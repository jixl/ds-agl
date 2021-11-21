package leetcode

// import "fmt"

// 567. 字符串的排列
// 给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
// 换句话说，s1 的排列之一是 s2 的 子串
// 输入：s1 = "ab" s2 = "eidbaooo"
// 输出：true
// 解释：s2 包含 s1 的排列之一 ("ba").
func checkInclusion(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)
	if m > n {
		return false
	}
	bs, ck := [26]int{}, [26]int{}
	for _, v := range s1 {
		bs[v-'a'] += 1
	}
	for _, v := range s2[0 : m-1] {
		ck[v-'a'] += 1
	}
	var find bool
	for i, v := range s2[m-1:] {
		ck[v-'a'] += 1

		for _, v1 := range s2[i : i+m] {
			idx := v1 - 'a'
			if ck[idx] != bs[idx] {
				find = false
				break
			}
			find = true
		}
		ck[s2[i]-'a'] -= 1
		if find {
			return find
		}
	}
	return find
}
