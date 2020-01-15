package q2

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestQuestion(t *testing.T) {
	l1 := &ListNode{
		Val:  9,
		Next: nil,
	}
	l1_1 := &ListNode{
		Val:  9,
		Next: nil,
	}
	l1.Next = l1_1


	l2 := &ListNode{
		Val:  9,
		Next: nil,
	}

	result := addTwoNumbers(l1, l2)

	c := result

	for {
		t.Log(c.Val)
		if c.Next == nil {
			break
		}
		c = c.Next
	}

}

func TestQuestion2(t *testing.T) {
	l1 := &ListNode{
		Val:  9,
		Next: nil,
	}
	l1_1 := &ListNode{
		Val:  9,
		Next: nil,
	}
	l1.Next = l1_1


	l2 := &ListNode{
		Val:  9,
		Next: nil,
	}

	result := addTwoNumbers2(l1, l2)

	c := result

	for {
		t.Log(c.Val)
		if c.Next == nil {
			break
		}
		c = c.Next
	}

}

// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	currentL1 := l1
	currentParentL1 := l1
	currentL2 := l2

	upNumber := 0

	for ;currentL1 != nil || currentL2 != nil; {

		if currentL1 == nil {
			currentParentL1.Next = new(ListNode)
			currentL1 = currentParentL1.Next
			currentParentL1 = currentL1
		}

		if currentL2 != nil {
			currentL1.Val += currentL2.Val
			currentL2 = currentL2.Next
		}

		currentL1.Val += upNumber
		upNumber = currentL1.Val / 10
		currentL1.Val %= 10

		currentParentL1 = currentL1
		currentL1 = currentL1.Next
	}

	if upNumber > 0 {
		currentParentL1.Next = new(ListNode)
		currentL1 = currentParentL1.Next
		currentL1.Val = upNumber
	}

	return l1
}

//
// 下面是另外一种解法
//

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {

	var headNode *ListNode
	var tmpNode *ListNode
	var upNum int = 0

	if l1 == nil || l2 == nil {
		return headNode
	}

	ct1 := l1
	ct2 := l2

	for {
		currentNode := new(ListNode)
		currentNode.Val = GetNodeValue(ct1) + GetNodeValue(ct2) + upNum
		if upNum > 0 {
			upNum--
		}

		if currentNode.Val > 9 {
			currentNode.Val -= 10
			upNum++
		}

		if headNode == nil  {
			headNode = currentNode
			tmpNode = currentNode
		} else {
			tmpNode.Next = currentNode
			tmpNode = currentNode
		}

		if (ct1 == nil || ct1.Next == nil) && (ct2 == nil || ct2.Next == nil) {
			break
		}

		ct1 = GetNodeNext(ct1)
		ct2 = GetNodeNext(ct2)
	}

	if upNum > 0 {
		currentNode := new(ListNode)
		currentNode.Val = upNum
		tmpNode.Next = currentNode
		tmpNode = currentNode
	}

	return headNode
}

func GetNodeValue(node *ListNode) int {
	if  node == nil {
		return 0
	} else {
		return node.Val
	}
}

func GetNodeNext(node *ListNode) *ListNode {
	if  node == nil {
		return nil
	} else {
		return node.Next
	}
}
