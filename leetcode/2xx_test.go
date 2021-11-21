package leetcode

import (
	"reflect"
	"strconv"
	"testing"
)

var twoSumCase = []struct {
	inputs []int
	target int
	want   []int
}{
	{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
	{[]int{2, 3, 4}, 6, []int{1, 3}},
	{[]int{-1, 0}, -1, []int{1, 2}},
	{[]int{5, 25, 75}, 100, []int{2, 3}},
}

func TestTwoSum(t *testing.T) {
	for idx, test := range twoSumCase {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := twoSum(test.inputs, test.target)
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("失败！got:%#v want:%#v\n", got, test.want)
			}
		})
	}
}

var containsDuplicateCase = []struct {
	inputs []int
	want   bool
}{
	{[]int{1, 2, 3, 1}, true},
	{[]int{1, 2, 3, 4}, false},
	{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
}

func TestContainsDuplicate(t *testing.T) {
	for idx, test := range containsDuplicateCase {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := containsDuplicate(test.inputs)
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("失败！got:%#v want:%#v\n", got, test.want)
			}
		})
	}
}

var moveZeroesCase = []struct{ inputs, want []int }{
	{[]int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}, []int{4, 2, 4, 3, 5, 1, 0, 0, 0, 0}},
	{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	{[]int{3, 2}, []int{3, 2}},
	{[]int{0, 0}, []int{0, 0}},
	{[]int{1, 0}, []int{1, 0}},
	{[]int{0}, []int{0}},
}

func TestMoveZeroes(t *testing.T) {
	for idx, test := range moveZeroesCase {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := moveZeroes(test.inputs)
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("失败！got:%#v want:%#v\n", got, test.want)
			}
		})
	}
}

var matrixReshapeCase = []struct {
	inputs [][]int
	r, c   int
	want   [][]int
}{
	{[][]int{{1, 2}, {3, 4}}, 1, 4, [][]int{{1, 2, 3, 4}}},
	{[][]int{{1, 2}, {3, 4}}, 2, 4, [][]int{{1, 2}, {3, 4}}},
}

func TestMatrixReshape(t *testing.T) {
	for idx, test := range matrixReshapeCase {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := matrixReshape(test.inputs, test.r, test.c)
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("失败！got:%#v want:%#v\n", got, test.want)
			}
		})
	}
}

var generateCase = []struct {
	rows int
	want [][]int
}{
	{1, [][]int{{1}}},
	{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
}

func TestGenerate(t *testing.T) {
	for idx, test := range generateCase {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := generate(test.rows)
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("失败！got:%#v want:%#v\n", got, test.want)
			}
		})
	}
}

var isAnagramCase = []struct {
	in1, in2 string
	want     bool
}{
	{"anagram", "nagaram", true},
	{"rat", "cat", false},
}

func TestIsAnagram(t *testing.T) {
	for idx, test := range isAnagramCase {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := isAnagram(test.in1, test.in2)
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("失败！got:%#v want:%#v\n", got, test.want)
			}
		})
	}
}
