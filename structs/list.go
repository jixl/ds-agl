package structs

type List interface {
	Insert(elem interface{}, pos int)
	Delete(elem interface{})
	Push(elem interface{})
	Pop()
	IsEmpty()
	Find(elem interface{})
	FindPrevious(elem interface{})
	Union(list *List) List
	Print()
}

type LinkedNode struct {
	value interface{}
	next  *LinkedNode
}

type DoubleNode struct {
	value      interface{}
	prev, next *DoubleNode
}

// func (list *List) Insert(data int) {
// 	head := list.head
// 	node := Node{data: data}
// 	if head == nil {
// 		list.head = &node
// 		list.tail = &node
// 	} else if head.data >= node.data {
// 		node.next = head
// 		list.head = &node
// 	} else {
// 		for head.data < node.data {
// 			if head.next == nil {
// 				head.next = &node
// 				list.tail = &node
// 				break
// 			} else if head.next.data >= node.data {
// 				node.next = head.next
// 				head.next = &node
// 				break
// 			} else {
// 				head = head.next
// 			}
// 		}
// 	}
// 	list.size += 1
// }

// func (list *List) Remove(data int) bool {
// 	node := list.head
// 	if node == nil {
// 		return false
// 	} else if node.data == data {
// 		node = node.next
// 		list.size -= 1
// 		if node == nil {
// 			list.tail = nil
// 		}
// 		return true
// 	} else {
// 		for node.next != nil {
// 			tmp := node.next
// 			if tmp.data == data {
// 				if tmp == list.tail {
// 					list.tail = node
// 				}
// 				node.next = tmp.next
// 				list.size -= 1
// 				return true
// 			} else {
// 				node = tmp
// 			}
// 		}
// 	}
// 	return false
// }

// func (list *List) Size() int {
// 	return list.size
// }

// func (list *List) Print() {
// 	node := list.head
// 	fmt.Printf("Size=%d\n", list.size)
// 	for node != nil {
// 		fmt.Print(node.data, " ")
// 		node = node.next
// 	}
// 	fmt.Print("\n")
// }
