package main

import (
	"fmt"
)

type LNode struct {
	Data interface{}
	Next *LNode
}

func CreateNode(node *LNode, max int) {
	cur := node
	for i := 1; i < max; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = i
		cur = cur.Next
	}
}

func PrintNode(info string, node *LNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}

//单链表——逆序
func rev(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}
	var pre *LNode
	var cur *LNode
	next := node.Next

	//next = 1 cur = 2 next.Next = nil pre = 1  next = 2
	//cur = 3 next.Next = 1 pre = 2 next = 3
	//cur = 4 next.Next = 2 pre = 3 next = 4
	//cur = nil next.Next = 3 pre = 4 next = nil
	//node.Next = 4

	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}

	node.Next = pre

}

//单链表——插入

func Ine(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}

	var cur *LNode
	var next *LNode
	cur = node.Next.Next
	node.Next.Next = nil
	//cur = 2 next = 3 cur.Next = 1 node.Next = 2 cur = 3
	//next = 4 cur.Next = 2  node.Next = 3 cur = 4
	//next = nil cur.Next = 3 node.Next = 4 cur = nil
	for cur != nil {
		next = cur.Next
		fmt.Println(next)
		cur.Next = node.Next
		fmt.Println(cur.Next)
		node.Next = cur
		fmt.Println(node.Next)
		cur = next
		fmt.Println(cur)
	}

}

//无序链表重复项
func rem(head *LNode) {
	if head == nil || head.Next == nil {
		return
	}

	outerCur := head.Next

	for ; outerCur != nil; outerCur = outerCur.Next {
		for innerCur, innerPre := outerCur.Next, outerCur; innerCur != nil; {
			if outerCur.Data == innerCur.Data {
				innerPre.Next = innerCur.Next
				innerCur = innerCur.Next
			} else {
				innerPre = innerCur
				innerCur = innerCur.Next
			}
		}
	}

}

//链表数之和
func add(l1 *LNode, l2 *LNode) *LNode {
	if l1 == nil || l1.Next == nil {
		return l2
	}
	if l2 == nil || l2.Next == nil {
		return l1
	}

	n1 := l1.Next
	n2 := l2.Next
	c := 0
	result := &LNode{}
	p := result
	for n1 != nil && n2 != nil {
		p.Next = &LNode{}
		sum := n1.Data.(int) + n2.Data.(int) + c
		p.Next.Data = sum % 10
		c = sum / 10
		p = p.Next
		n1 = n1.Next
		n2 = n2.Next
	}
	if n1 == nil {
		for n2 != nil {
			sum := n2.Data.(int) + c
			p.Next = &LNode{}
			p.Next.Data = sum % 10
			c = sum / 10
			p = p.Next
			n2 = n2.Next
		}
	}

	if n2 == nil {
		for n1 != nil {
			sum := n1.Data.(int) + c
			p.Next = &LNode{}
			p.Next.Data = sum % 10
			c = sum / 10
			p = p.Next
			n1 = n1.Next
		}
	}

	if c == 1 {
		p.Next = &LNode{}
		p.Next.Data = 1
	}

	return result
}

func createNodeT() (l1 *LNode, l2 *LNode) {
	l1 = &LNode{}
	l2 = &LNode{}
	cur := l1
	for i := 1; i < 7; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = i + 2
		cur = cur.Next
	}

	cur = l2
	for i := 9; i > 4; i-- {
		cur.Next = &LNode{}
		cur.Next.Data = i
		cur = cur.Next
	}

	return

}

//中间节点
func findMidNode(node *LNode) *LNode {
	if node == nil || node.Next == nil {
		return node
	}

	fast := node
	slow := node
	slowPre := node

	for fast != nil && fast.Next != nil {
		slowPre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	slowPre.Next = nil
	return slow
}

//无head结点链表逆序
func reverseL(node *LNode) *LNode {
	if node == nil || node.Next == nil {
		return node
	}

	var pre *LNode

	//next=5 node.Next = nil pre = 4 node = 5
	//next = 6 node.Next = 4 pre = 5 node = 6
	//next = 7 node.Next = 5 pre = 6 node = 7
	//next = nil node.Next = 6 pre = 7 node = nil
	for node != nil {
		next := node.Next
		node.Next = pre
		pre = node
		node = next
	}

	return pre

}

func reorder(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}

	cur1 := node.Next

	mid := findMidNode(node.Next)
	cur2 := reverseL(mid)
	var tmp *LNode

	for cur1.Next != nil {
		tmp = cur1.Next
		cur1.Next = cur2
		cur1 = tmp
		tmp = cur2.Next
		cur2.Next = cur1
		cur2 = tmp
	}
	cur1.Next = cur2
}

//指针法k
func findLastK(node *LNode, k int) *LNode {
	if node == nil || node.Next == nil {
		return node
	}

	slow := node.Next
	fast := node.Next
	i := 0
	for i = 0; i < k && fast != nil; i++ {
		fast = fast.Next
	}
	if i < k {
		return nil
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	return slow

}

func rotateK(node *LNode, k int) {
	if node == nil || node.Next == nil {
		return
	}

	slow := node
	fast := node
	i := 0

	for i = 0; i < k && fast != nil; i++ {
		fast = fast.Next
	}
	if i < k {
		return
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	tmp := slow
	slow = slow.Next
	tmp.Next = nil
	fast.Next = node.Next
	node.Next = slow

}

//判断链表是否有重复数环
func isLoop(node *LNode) *LNode {
	if node == nil || node.Next == nil {
		return node
	}

	slow := node.Next
	fast := node.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return slow
		}
	}
	return nil

}

func findLoopNode(node *LNode, meetNode *LNode) *LNode {
	first := node.Next
	second := meetNode

	for first != second {
		first = first.Next
		second = second.Next
	}
	return first

}

//相邻元素
func reverseN(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}

	cur := node.Next
	pre := node
	var next *LNode
	//next = 3 pre.Next = 2 cur.Next.Next = 1 cur.Next = 3 pre = 1 cur = 3
	//next = 5 pre.Next = 4 cur.Next.Next = 3 cur.Next = 5 pre = 3 cur = 5
	//next = 7 pre.Next = 6 cur.Next.Next = 5 cur.Next = 7 pre = 5 cur = 7
	for cur != nil && cur.Next != nil {
		next = cur.Next.Next
		pre.Next = cur.Next
		cur.Next.Next = cur
		cur.Next = next
		pre = cur
		cur = next

	}

}

//k个结点为一组
func reverseK(node *LNode, k int) {
	if node == nil || node.Next == nil {
		return
	}
	pre := node
	begin := node.Next
	for begin != nil {
		end := begin

		for i := 1; i < k; i++ {
			if end.Next != nil {
				end = end.Next
			} else {
				return
			}
		}
		tmp := end.Next
		end.Next = nil
		pre.Next = reverseL(begin)
		begin.Next = tmp
		pre = begin
		begin = tmp

	}

}

//两个有序链表
func merge(nodeOne *LNode, nodeTwo *LNode) *LNode {
	if nodeOne == nil || nodeOne.Next == nil {
		return nodeOne
	}
	if nodeTwo == nil || nodeTwo.Next == nil {
		return nodeTwo
	}

	headOne := nodeOne.Next
	headTwo := nodeTwo.Next
	var head *LNode
	var cur *LNode
	if headOne.Data.(int) > headTwo.Data.(int) {
		head = nodeTwo
		cur = headTwo
		headTwo = headTwo.Next
	} else {
		head = nodeOne
		cur = headOne
		headOne = headOne.Next
	}
	for headOne != nil && headTwo != nil {
		if headOne.Data.(int) < headTwo.Data.(int) {
			cur.Next = headOne
			cur = headOne
			headOne = headOne.Next
		} else {
			cur.Next = headTwo
			cur = headTwo
			headTwo = headTwo.Next
		}
	}

	if headOne != nil {
		cur.Next = headOne
	}
	if headTwo != nil {
		cur.Next = headTwo
	}
	return head

}

func createNodeN(node *LNode, start int) {
	cur := node
	for i := start; i < 7; i += 2 {
		cur.Next = &LNode{}
		cur.Next.Data = i
		cur = cur.Next
	}
}

//指定结点
func removeNode(node *LNode) bool {
	if node == nil || node.Next == nil {
		return false
	}
	node.Data = node.Next.Data
	tmp := node.Next
	node.Next = tmp.Next
	return true
}

func createNodeNr(Node *LNode, n int) (retNode *LNode) {
	cur := Node
	for i := 1; i < 8; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = i
		cur = cur.Next
		if i == n {
			retNode = cur
		}
	}
	return
}

//两个链表是否重复
func isIntersect(nodeOne *LNode, nodeTwo *LNode) *LNode {
	if nodeOne == nil || nodeOne.Next == nil || nodeTwo == nil || nodeTwo.Next == nil || nodeOne == nodeTwo {
		return nil
	}
	curOne := nodeOne
	curTwo := nodeTwo
	nOne := 0
	nTwo := 0
	for curOne.Next != nil {
		nOne++
		curOne = curOne.Next
	}
	for curTwo.Next != nil {
		nTwo++
		curTwo = curTwo.Next
	}

	if curOne == curTwo {
		if nOne > nTwo {
			for nOne-nTwo > 0 {
				nodeOne = nodeOne.Next
				nOne--
			}
		}
		if nTwo > nOne {
			for nTwo-nOne > 0 {
				nodeTwo = nodeTwo.Next
				nTwo--
			}
		}

	}

	for nodeOne != nodeTwo {
		nodeOne = nodeOne.Next
		nodeTwo = nodeTwo.Next
	}

	return nodeOne

}

type LNodeN struct {
	data  int
	right *LNodeN
	down  *LNodeN
}

func (p *LNodeN) Insert(headRef *LNodeN, data int) *LNodeN {
	newNode := &LNodeN{
		data: data,
		down: headRef,
	}
	headRef = newNode
	return headRef
}

//链表排序
func mergeN(a *LNodeN, b *LNodeN) *LNodeN {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	var result *LNodeN

	if a.data < b.data {
		result = a
		result.down = mergeN(a.down, b)

	} else {
		result = b
		result.down = mergeN(a, b.down)

	}

	return result
}
func flatten(root *LNodeN) *LNodeN {
	if root == nil || root.right == nil {
		return root
	}

	root.right = flatten(root.right)
	root = mergeN(root, root.right)
	return root
}

func CreateNodeN() *LNodeN {
	node := &LNodeN{}
	node = node.Insert(nil, 31)
	node = node.Insert(node, 8)
	node = node.Insert(node, 6)
	node = node.Insert(node, 3)

	node.right = node.Insert(node.right, 21)
	node.right = node.Insert(node.right, 11)

	node.right.right = node.Insert(node.right.right, 50)
	node.right.right = node.Insert(node.right.right, 22)
	node.right.right = node.Insert(node.right.right, 15)

	node.right.right.right = node.Insert(node.right.right.right, 55)
	node.right.right.right = node.Insert(node.right.right.right, 40)
	node.right.right.right = node.Insert(node.right.right.right, 39)
	node.right.right.right = node.Insert(node.right.right.right, 30)

	return node

}

func PrintNodeN(info string, node *LNodeN) {
	fmt.Print(info)
	tmp := node
	for tmp != nil {
		fmt.Print(tmp.data, " ")
		tmp = tmp.down
	}
	fmt.Println()
}
func main() {

}
