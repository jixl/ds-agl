package structs

import (
	"reflect"
	"testing"
)

func TestLinkedListPush(t *testing.T) {
	list := new(LinkedList)
	list.Push(3)
	list.Push(1)
	list.Push(2)
	list.Push("23")
	list.Print()
	size := 4
	if !reflect.DeepEqual(size, list.Size()) {
		t.Fatalf("Push失败！got:%#v want:%#v\n", size, list.Size())
	}
}

func TestLinkedListInsert(t *testing.T) {
	list := new(LinkedList)
	list.Insert(3, 0)
	list.Print()
	list.Insert(1, 0)
	list.Print()
	list.Insert(2, 0)
	list.Print()
	list.Insert("23", 1)
	list.Print()
	size := 4
	if !reflect.DeepEqual(size, list.Size()) {
		t.Fatalf("Insert失败！got:%#v want:%#v\n", size, list.Size())
	}
}

func TestLinkedListIsEmpty(t *testing.T) {
	list := new(LinkedList)
	if !reflect.DeepEqual(true, list.IsEmpty()) {
		t.Fatalf("IsEmpty失败！got:%#v want:%#v\n", true, list.IsEmpty())
	}
	list.Insert("23", 2)
	list.Print()
	if !reflect.DeepEqual(false, list.IsEmpty()) {
		t.Fatalf("IsEmpty失败！got:%#v want:%#v\n", false, list.IsEmpty())
	}
}

func TestLinkedListDelete(t *testing.T) {
	list := new(LinkedList)
	list.Push(3)
	list.Push(1)
	list.Push(2)
	list.Push("23")
	list.Print()
	size := 4
	if !reflect.DeepEqual(size, list.Size()) {
		t.Fatalf("Insert失败！got:%#v want:%#v\n", size, list.Size())
	}

	list.Remove(2)
	list.Print()
	size = 3
	if !reflect.DeepEqual(size, list.Size()) {
		t.Fatalf("Insert失败！got:%#v want:%#v\n", size, list.Size())
	}
}

func TestLinkedListUnion(t *testing.T) {
	want := new(LinkedList)
	list := new(LinkedList)
	for i := 1; i < 10; i += 2 {
		list.Push(i)
		want.Push(i)
	}

	list2 := new(LinkedList)
	for i := 2; i < 10; i += 2 {
		list2.Push(i)
		want.Push(i)
	}
	ulist := list.Union(list2)
	ulist.Print()
	if !reflect.DeepEqual(ulist, want) {
		t.Fatalf("Union失败！got:%#v want:%#v\n", ulist, want)
	}
}

func TestLinkedListIntersect(t *testing.T) {
	list := new(LinkedList)
	for i := 1; i < 10; i += 2 {
		list.Push(i)
	}

	want := new(LinkedList)
	list2 := new(LinkedList)
	for i := 5; i < 10; i += 2 {
		list2.Push(i)
		if i < 10 {
			want.Push(i)
		}
	}
	ulist := list.Intersect(list2)
	ulist.Print()
	if !reflect.DeepEqual(ulist, want) {
		t.Fatalf("Intersect失败！got:%#v want:%#v\n", ulist, want)
	}
}
