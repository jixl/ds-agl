package structs

import "fmt"

type LinkedList struct {
	size       int
	head, tail *LinkedNode
}

func (list *LinkedList) Push(ele interface{}) {
	elem := LinkedNode{value: ele}
	if list.IsEmpty() {
		list.head = &LinkedNode{next: &elem}
		list.tail = &elem
	} else {
		list.tail.next = &elem
		list.tail = &elem
	}
	list.size++
}

func (list *LinkedList) Insert(ele interface{}, pos int) {
	if list.IsEmpty() || pos >= list.size {
		list.Push(ele)
	} else {
		node := list.posPrevious(pos)
		elem := LinkedNode{value: ele, next: node.next}
		node.next = &elem
		list.size++
	}
}

func (list *LinkedList) posPrevious(pos int) *LinkedNode {
	node := list.head
	for idx := 0; node.next != nil && idx < pos; idx++ {
		node = node.next
	}
	return node
}

func (list *LinkedList) findPrevious(ele interface{}) *LinkedNode {
	node := list.head
	for node.next != nil && node.next.value != ele {
		node = node.next
	}
	return node
}

func (list *LinkedList) Find(ele interface{}) bool {
	node := list.head.next
	for node != nil && node.value != ele {
		node = node.next
	}
	return node != nil
}

func (list *LinkedList) Union(ll *LinkedList) *LinkedList {
	if ll.IsEmpty() && list.IsEmpty() {
		return nil
	}

	unionList := &LinkedList{}
	node := list.head.next
	for node != nil {
		unionList.Push(node.value)
		node = node.next
	}
	node = ll.head.next
	for node != nil {
		unionList.Push(node.value)
		node = node.next
	}
	return unionList
}

func (list *LinkedList) Intersect(ll *LinkedList) *LinkedList {
	if list.IsEmpty() || ll.IsEmpty() {
		return nil
	}
	var node *LinkedNode
	var fList *LinkedList
	if list.Size() >= ll.Size() {
		node = ll.head.next
		fList = list
	} else {
		node = list.head.next
		fList = ll
	}

	inList := &LinkedList{}
	for node != nil {
		if fList.Find(node.value) {
			inList.Push(node.value)
		}
		node = node.next
	}
	return inList
}

func (list *LinkedList) Pop() interface{} {
	node := list.head
	if node.next != nil {
		tmp := node.next
		node.next = tmp.next
		node = tmp
	}
	return node.value
}

func (list *LinkedList) Remove(ele interface{}) {
	node := list.findPrevious(ele)
	if node.next != nil {
		if list.tail == node.next.next {
			list.tail = node
		}
		node.next = node.next.next
		list.size--
	}
}

func (list *LinkedList) Clear() {
	list.head.next = nil
	list.tail = nil
	list.size = 0
}

func (list *LinkedList) IsEmpty() bool {
	return list.head == nil
}

func (list *LinkedList) Size() int {
	return list.size
}

func (list *LinkedList) Print() {
	node := list.head
	fmt.Printf("Size=%d\n", list.size)
	for node.next != nil {
		fmt.Print(node.next.value, " ")
		node = node.next
	}
	fmt.Print("\n")
}
