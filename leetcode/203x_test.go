package leetcode

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

var twoOutOfThreeCase = []struct {
	inputs [][]int
	want   []int
}{
	{[][]int{{1, 1, 3, 2}, {2, 3}, {3}}, []int{2, 3}},
	{[][]int{{3, 1}, {2, 3}, {1, 2}}, []int{1, 2, 3}},
}

func TestTwoOutOfThree(t *testing.T) {
	for idx, test := range twoOutOfThreeCase {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := twoOutOfThree(test.inputs[0], test.inputs[1], test.inputs[2])
			sort.Ints(got)
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("失败！got:%#v want:%#v\n", got, test.want)
			}
		})
	}
}

var minOperationsCase = []struct {
	grid    [][]int
	x, want int
}{
	{[][]int{{2, 4}, {6, 8}}, 2, 4},
	{[][]int{{3, 1}, {2, 5}}, 1, 5},
	{[][]int{{3, 1}, {2, 5}}, 2, -1},
}

func TestMinOperations(t *testing.T) {
	for idx, test := range minOperationsCase {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := minOperations(test.grid, test.x)
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("失败！got:%#v want:%#v\n", got, test.want)
			}
		})
	}
}

func TestStockPrices(t *testing.T) {
}
