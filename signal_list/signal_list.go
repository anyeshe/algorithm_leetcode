package main

import "fmt"

type Node struct {
	value string
	next *Node
}

type SignalList struct {
	len int
	head *Node
	tail *Node
}

func New() *SignalList {
	return &SignalList{
		len:  0,
		head: nil,
		tail: nil,
	}
}

func (list *SignalList) Append(node *Node) bool {
	if node == nil {
		return false
	}

	node.next = nil

	if list.len == 0 {
		list.head = node
	} else {
		list.tail.next = node
	}

	list.tail = node
	list.len++

	return true
}

func (list *SignalList) Shift(node *Node) bool {
	if node == nil {
		return false
	}

	node.next = list.head
	list.head = node
	list.len++
	return true
}

func (list *SignalList) Insert(node *Node, index int) bool {
	if node == nil || list.len < index || list.len == 0 {
		return false
	}

	if index == 0 {
		node.next = list.head
		list.head = node
	} else {
		tmpNode := list.head
		for i := 1; i < index; i++ {
			tmpNode = tmpNode.next
		}
		node.next = tmpNode.next
		tmpNode.next = node
	}

	list.len++
	return true
}

func (list *SignalList) Show() {
	var tmpNode *Node = list.head
	for {
		if tmpNode == nil {
			break
		}
		fmt.Println(tmpNode.value)
		tmpNode = tmpNode.next
	}
}

func (list *SignalList) delete(index int) bool {
	if index > list.len {
		return false
	}

	if index == 0 {
		newHead := list.head.next
		list.head = newHead
		if list.len == 1 {
			list.tail = newHead
		}
	} else {
		preNode := list.head
		// 不能遍历到要删除的节点上, 那有就找不到前面的了
		for i := 1; i < index; i++ {
			preNode = preNode.next
		}

		tmpNode := preNode.next
		preNode.next = tmpNode.next

		if index == list.len-1 {
			list.tail = preNode
		}
	}

	list.len--
	return true
}

func (list *SignalList) Reverse() bool {
	if list.len <= 1 {
		return true
	}

	currentNode := list.head
	var proNode *Node = nil
	var nextNode *Node = nil

	for currentNode != nil {
		// nextNode指向下一个节点 - 防止丢了
		// 当前节点的next指向上一个节点 - 反转
		// ----------  为下一次准备  ----------
		// 上一个节点指向当前节点
		// 当前节点指向下一个节点
		nextNode = currentNode.next
		currentNode.next = proNode
		proNode = currentNode
		currentNode = nextNode
	}

	list.head, list.tail = list.tail, list.head

	return true
}

func main() {
	list := New()

	list.Append(&Node{
		value: "a",
		next:  nil,
	})

	list.Append(&Node{
		value: "b",
		next:  nil,
	})
	
	list.Shift(&Node{
		value: "s1",
		next:  nil,
	})

	list.Insert(&Node{
		value: "I1",
		next:  nil,
	}, 2)

	list.delete(2)

	list.Show()

	list.Reverse()

	list.Show()
}