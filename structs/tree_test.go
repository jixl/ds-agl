package structs

import (
	"reflect"
	"strconv"
	"testing"
)

/**
	Tree
    3
  /   \
 9     20
  \    /  \
  10  15   7
**/

type treeOrderCase struct {
	inputs *TreeNode
	want   []int
}

var case1 = []interface{}{3, 9, 20, nil, 10, 15, 7}

var preOrderCase = []treeOrderCase{
	{NewTree(case1), []int{3, 9, 10, 20, 15, 7}},
}

func TestPreOrder(t *testing.T) {
	for idx, test := range preOrderCase {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := test.inputs.PreOrder()
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("失败！got:%#v want:%#v\n", got, test.want)
			}
		})
	}
}

var inOrderCase = []treeOrderCase{
	{NewTree(case1), []int{9, 10, 3, 15, 20, 7}},
}

func TestInOrder(t *testing.T) {
	for idx, test := range inOrderCase {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := test.inputs.InOrder()
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("失败！got:%#v want:%#v\n", got, test.want)
			}
		})
	}
}

var postOrderCase = []treeOrderCase{
	{NewTree(case1), []int{10, 9, 15, 7, 20, 3}},
}

func TestPostOrder(t *testing.T) {
	for idx, test := range postOrderCase {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := test.inputs.PostOrder()
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("失败！got:%#v want:%#v\n", got, test.want)
			}
		})
	}
}

var levelOrderCase = []treeOrderCase{
	{NewTree(case1), []int{3, 9, 20, 10, 15, 7}},
}

func TestLevelOrder(t *testing.T) {
	for idx, test := range levelOrderCase {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := test.inputs.LevelOrder()
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("失败！got:%#v want:%#v\n", got, test.want)
			}
		})
	}
}
