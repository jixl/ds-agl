// Package morestrings implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package leetcode

import (
	"fmt"
	// "unsafe"
	"strings"
	"unicode/utf8"
)

// ReverseRunes returns its argument string reversed rune-wise left to right.
func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func reverseString(s []byte) {
	for l, r := 0, len(s)-1; l < r; {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

func reverseWords(s string) string {
	i, bs := 0, []rune(s)
	for _, w := range strings.Split(s, " ") {
		if len(w) != 0 {
			for l, r := i, i+utf8.RuneCountInString(w)-1; l < r; l, r = l+1, r-1 {
				bs[l], bs[r] = bs[r], bs[l]
			}
			i += len(w)
		}
		i++
		fmt.Printf("%d %d %d\n", i, len(w), len(bs))
	}
	return string(bs)
}

func loc(i, j, v byte) string {
	var x, y byte
	x, y = (i+3)/3, (j+3)/3
	a := string(x*y + 10 + v)
	// a := (*string)(unsafe.Pointer(&[]byte{x*y, v}))
	// fmt.Printf(*a)
	fmt.Printf("%d %d %d %v xxx", x, y, v, a)
	return a
}
