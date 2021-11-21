package leetcode

// import "fmt"

// 387. 字符串中的第一个唯一字符
// 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
func firstUniqChar(s string) int {
	r := []byte(s)
	mp := make(map[byte]int)
	for i, v := range r {
		if _, ok := mp[v]; ok {
			mp[v] = -1
		} else {
			mp[v] = i
		}
	}
	for i, v := range r {
		if mp[v] != -1 {
			return i
		}
	}
	return -1
}

func firstUniqChar1(s string) int {
	r := [26]int{}
	for i := 0; i < len(s); i++ {
		r[s[i]-'a'] = r[s[i]-'a'] + 1
	}

	for i := 0; i < len(s); i++ {
		if r[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

// 383. 赎金信 https://leetcode-cn.com/problems/ransom-note/
func canConstruct(ransomNote string, magazine string) bool {
	m, n := len(ransomNote), len(magazine)
	if m > n {
		return false
	}

	bs := [26]int{}
	for i := 0; i < n; i++ {
		bs[magazine[i]-'a'] = bs[magazine[i]-'a'] + 1
		if i < m {
			bs[ransomNote[i]-'a'] = bs[ransomNote[i]-'a'] - 1
		}
	}

	for i := 0; i < m; i++ {
		if bs[ransomNote[i]-'a'] < 0 {
			return false
		}
	}

	return true
}
