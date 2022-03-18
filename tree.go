package main

import (
	"fmt"
	"math"
	"sync"
)

type BNode struct {
	data       interface{}
	leftChild  *BNode
	rightChild *BNode
}

func newBNode() *BNode {
	return &BNode{
		data:       nil,
		leftChild:  nil,
		rightChild: nil,
	}
}

//数组二叉树
func arrayToTree(arr []int, start int, end int) *BNode {
	var root *BNode
	if end >= start {
		root = newBNode()
		mid := (start + end + 1) / 2
		root.data = arr[mid]
		root.leftChild = arrayToTree(arr, start, mid-1)
		root.rightChild = arrayToTree(arr, mid+1, end)
	}
	return root
}

func printTreeMidOrder(root *BNode) {
	if root == nil {
		return
	}
	if root.leftChild != nil {
		printTreeMidOrder(root.leftChild)
	}
	fmt.Print(root.data, " ")
	if root.rightChild != nil {
		printTreeMidOrder(root.rightChild)
	}
}

type sliceQueueTree struct {
	arr []interface{}
	sync.RWMutex
}

func (p *sliceQueueTree) isEmpty() bool {
	return p.size() == 0
}

func (p *sliceQueueTree) size() int {
	return len(p.arr)
}
func (p *sliceQueueTree) enQueue(data interface{}) {
	p.Lock()
	defer p.Unlock()
	p.arr = append(p.arr, data)
}

func (p *sliceQueueTree) deQueue() interface{} {
	p.Lock()
	defer p.Unlock()
	if !p.isEmpty() {
		ret := p.arr[0]
		p.arr = p.arr[1:]
		return ret
	}
	return nil
}

func (p *sliceQueueTree) popQueue() interface{} {
	p.Lock()
	defer p.Unlock()
	if !p.isEmpty() {
		ret := p.arr[p.size()-1]
		p.arr = p.arr[:p.size()-1]
		return ret
	}
	return nil
}
func (p *sliceQueueTree) getFront() interface{} {
	p.RLock()
	defer p.RUnlock()
	if !p.isEmpty() {
		return p.arr[0]
	}
	return nil
}

func (p *sliceQueueTree) getBack() interface{} {
	p.RLock()
	defer p.RUnlock()
	if !p.isEmpty() {
		return p.arr[p.size()-1]
	}
	return nil
}

func (p *sliceQueueTree) enQueueF(data interface{}) {
	p.Lock()
	defer p.Unlock()
	p.arr = append([]interface{}{data}, p.arr...)
}
func (p *sliceQueueTree) remove(data interface{}) {
	p.Lock()
	defer p.Unlock()
	for k, v := range p.arr {
		if v == data {
			p.arr = append(p.arr[:k], p.arr[k+1:]...)
		}
	}
}

func (p *sliceQueueTree) list() interface{} {
	return p.arr
}

//二叉树结点
func printTree(root *BNode) {
	if root == nil {
		return
	}
	q := &sliceQueueTree{
		arr:     []interface{}{},
		RWMutex: sync.RWMutex{},
	}
	q.enQueue(root)
	for q.size() > 0 {
		r := q.deQueue().(*BNode)
		fmt.Print(r.data, " ")
		if r.leftChild != nil {
			q.enQueue(r.leftChild)
		}
		if r.rightChild != nil {
			q.enQueue(r.rightChild)
		}

	}

}

func printAtLevel(root *BNode, level int) int {
	if root == nil || level < 0 {
		return 0
	} else if level == 0 {
		fmt.Println(root.data)
		return 1
	} else {
		return printAtLevel(root.leftChild, level-1) + printAtLevel(root.rightChild, level-1)
	}

}

//二叉树子结点最大和
var maxSum = math.MinInt64

// root = 6 l = 3 l = -1  lMax = 0 rMax = 0 sum = -1
// root = 3 r = 9 lMax = 0 rMax = 0 sum = 9
// sum = 11
// root = 6 lMax = 11 r = -7 lMax = 0 rMax = 0 sum = -7
// root = 6 lMax = 11 r = -7 sum = 10
func findMaxTree(root *BNode, maxRoot *BNode) int {
	if root == nil {
		return 0
	}
	lMax := findMaxTree(root.leftChild, maxRoot)
	rMax := findMaxTree(root.rightChild, maxRoot)
	sum := lMax + rMax + root.data.(int)
	fmt.Println(lMax, rMax, root.data)
	if sum > maxSum {
		maxSum = sum
		maxRoot.data = root.data
	}

	return sum

}

func createTree() *BNode {
	root := &BNode{}
	node1 := &BNode{}
	node2 := &BNode{}
	node3 := &BNode{}
	node4 := &BNode{}
	root.data = 6
	node1.data = 3
	node2.data = -7
	node3.data = -1
	node4.data = 9
	root.leftChild = node1
	root.rightChild = node2
	node1.leftChild = node3
	node1.rightChild = node4
	return root

}

//二叉树判断相等
func isEqual(rootOne *BNode, rootTwo *BNode) bool {
	if rootOne == nil && rootTwo == nil {
		return true
	}
	if rootOne == nil && rootTwo != nil {
		return false
	}
	if rootOne != nil && rootTwo == nil {
		return false
	}
	if rootOne.data == rootTwo.data {
		return isEqual(rootOne.leftChild, rootTwo.leftChild) && isEqual(rootOne.rightChild, rootTwo.rightChild)
	}
	return false
}

var pHead *BNode
var pEnd *BNode

//二叉树双向链表
func inOrderTree(root *BNode) {
	if root == nil {
		return
	}

	inOrderTree(root.leftChild)
	root.leftChild = pEnd
	if pEnd == nil {
		pHead = root
	} else {
		pEnd.rightChild = root
	}
	pEnd = root
	inOrderTree(root.rightChild)
}

//是否二叉树遍历
func isAfterOrder(arr []int, start int, end int) bool {
	if arr == nil {
		return false
	}
	root := arr[end]
	var i, j int
	for i = start; i < end; i++ {
		if arr[i] > root {
			break
		}
	}
	for j = i; j < end; j++ {
		if arr[j] < root {
			return false
		}
	}

	l := true
	r := true

	if i > start {
		l = isAfterOrder(arr, start, i-1)
	}
	if j < end {
		r = isAfterOrder(arr, i, end)
	}

	return l && r

}

type sliceStackT struct {
	arr       []interface{}
	stackSize int
	sync.RWMutex
}

func (p *sliceStackT) isEmpty() bool {
	return p.stackSize == 0
}

func (p *sliceStackT) size() int {
	return p.stackSize
}
func (p *sliceStackT) push(t interface{}) {
	p.Lock()
	defer p.Unlock()
	p.arr = append(p.arr, t)
	p.stackSize = p.stackSize + 1
}
func (p *sliceStackT) pop() interface{} {
	if p.stackSize > 0 {
		p.stackSize--
		ret := p.arr[p.stackSize]
		p.arr = p.arr[:p.stackSize]
		return ret
	}
	return nil
}
func (p *sliceStackT) top() interface{} {
	if p.stackSize > 0 {
		return p.arr[p.stackSize-1]
	}
	return nil
}

func (p *sliceStackT) list() []interface{} {
	return p.arr
}

//路径对比
func getPathFromRoot(root *BNode, node *BNode, s *sliceStackT) bool {
	if root == nil {
		return false
	}

	if root.data == node.data {
		s.push(root)
		return true
	}

	if getPathFromRoot(root.leftChild, node, s) || getPathFromRoot(root.rightChild, node, s) {
		s.push(root)
		return true
	}
	return false
}

func findParentNode(root, node1, node2 *BNode) *BNode {
	sOne := &sliceStackT{}
	sTwo := &sliceStackT{}
	getPathFromRoot(root, node1, sOne)
	getPathFromRoot(root, node2, sTwo)
	var c *BNode
	for tOne, tTwo := sOne.pop().(*BNode), sTwo.pop().(*BNode); tOne != nil && tTwo != nil && tOne.data.(int) == tTwo.data.(int); {

		c = tOne
		tOne = sOne.pop().(*BNode)
		tTwo = sTwo.pop().(*BNode)
	}
	return c
}

//编号
func getNumber(root, node *BNode, number int) (bool, int) {
	if root == nil {
		return false, number
	}
	if root.data == node.data {
		return true, number
	}
	num := 2 * number
	var b bool
	if b, num = getNumber(root.leftChild, node, num); b {
		return true, num
	} else {
		num = 2*number + 1
		return getNumber(root.rightChild, node, num)
	}
}

func getNodeFromNum(root *BNode, number int) *BNode {
	if root == nil || number < 0 {
		return nil
	} else if number == 1 {
		return root
	}

	ll := (uint)(math.Log(float64(number)) / math.Log(2.0))
	number -= 1 << ll
	for ; ll > 0; ll-- {
		if (1 << (ll - 1) & number) == 1 {
			root = root.rightChild
		} else {
			root = root.leftChild
		}
	}
	return root
}

func findParentNodeNum(root, node1, node2 *BNode) *BNode {
	num1 := 1
	num2 := 1
	_, num1 = getNumber(root, node1, num1)
	_, num2 = getNumber(root, node2, num2)

	for num1 != num2 {
		if num1 > num2 {
			num1 /= 2
		} else {
			num2 /= 2
		}
	}
	return getNodeFromNum(root, num1)
}

//遍历
func findParentNodeReverse(root, node1, node2 *BNode) *BNode {
	if root == nil || root.data.(int) == node1.data.(int) || root.data.(int) == node2.data.(int) {
		return root
	}
	l := findParentNodeReverse(root.leftChild, node1, node2)
	r := findParentNodeReverse(root.rightChild, node1, node2)
	if l == nil {
		return r
	} else if r == nil {
		return l
	} else {
		return root
	}
}

//一次遍历
func lowestCommonAncestor(root, p, q *BNode) (ancestor *BNode) {
	ancestor = root
	for {
		if p.data.(int) < ancestor.data.(int) && q.data.(int) < ancestor.data.(int) {
			ancestor = ancestor.leftChild
		} else if p.data.(int) > ancestor.data.(int) && q.data.(int) > ancestor.data.(int) {
			ancestor = ancestor.rightChild
		} else {
			return
		}
	}
}

//二叉树
func dupTree(root *BNode) *BNode {
	if root == nil {
		return nil
	}

	var r *BNode
	r = root
	r.leftChild = dupTree(root.leftChild)
	r.rightChild = dupTree(root.rightChild)

	return r

}

//二叉树对应路径
func findRoad(root *BNode, num, sum int, v []int) {
	sum += root.data.(int)
	v = append(v, root.data.(int))
	if root.leftChild == nil && root.rightChild == nil && sum == num {
		for _, d := range v {
			fmt.Println(d, " ")
		}
		fmt.Println()
	}
	if root.leftChild != nil {
		findRoad(root.leftChild, num, sum, v)
	}
	if root.rightChild != nil {
		findRoad(root.rightChild, num, sum, v)
	}
	sum -= v[len(v)-1]
	v = v[:len(v)-1]

}

//二叉树镜像
func reverseTree(root *BNode) {
	if root == nil {
		return
	}
	reverseTree(root.leftChild)
	reverseTree(root.rightChild)
	tmp := root.leftChild
	root.leftChild = root.rightChild
	root.rightChild = tmp
}

//二叉排序树大于中间值结点查找
func getMinNode(root *BNode) *BNode {
	if root == nil {
		return root
	}
	tmp := root
	for tmp.leftChild != nil {
		tmp = tmp.leftChild
	}
	return tmp
}

func getMaxNode(root *BNode) *BNode {
	if root == nil {
		return root
	}
	tmp := root
	for tmp.rightChild != nil {
		tmp = tmp.rightChild
	}
	return tmp
}

func getNode(root *BNode) *BNode {
	min := getMinNode(root)
	max := getMaxNode(root)
	mid := (min.data.(int) + max.data.(int)) / 2
	var r *BNode
	for root != nil {
		if root.data.(int) <= mid {
			root = root.rightChild
		} else {
			r = root
			root = root.leftChild
		}
	}
	return r
}

type intRef struct {
	val int
}

func Max(a, b, c int) int {
	var max int
	if a < b {
		max = b
	} else {
		max = a
	}

	if max < c {
		max = c
	}
	return max
}

func findMaxPathRecursive(root *BNode, max *intRef) int {
	if root == nil {
		return 0
	}

	l := findMaxPathRecursive(root.leftChild, max)
	r := findMaxPathRecursive(root.rightChild, max)
	lMax := root.data.(int) + l
	rMax := root.data.(int) + r
	AllMax := l + r + root.data.(int)
	tmpMax := Max(AllMax, lMax, rMax)
	if tmpMax > max.val {
		max.val = tmpMax
	}
	var subMax int
	if l > r {
		subMax = l
	} else {
		subMax = r
	}
	return root.data.(int) + subMax

}

func findMaxPath(root *BNode) int {
	max := &intRef{val: math.MinInt64}
	findMaxPathRecursive(root, max)
	return max.val
}

type tireTree struct {
	isLeaf bool
	url    string
	child  []*tireTree
}

type dnsCache struct {
	root *tireTree
}

var charCount = 11

func (p *dnsCache) getIndexFromRune(char rune) int {
	if char == '.' {
		return 10
	} else {
		return int(char) - '0'
	}

}

func (p *dnsCache) getRuneFromIndex(i int) rune {
	if i == 10 {
		return '.'
	} else {
		return rune('0' + i)
	}
}

func (p *dnsCache) insert(ip, url string) {
	pCrawl := p.root
	for _, v := range []rune(ip) {
		index := p.getIndexFromRune(v)
		if pCrawl.child[index] == nil {
			pCrawl.child[index] = &tireTree{
				false,
				"",
				make([]*tireTree, charCount),
			}
		}
		pCrawl = pCrawl.child[index]
	}
	pCrawl.isLeaf = true
	pCrawl.url = url
}

//dns查找
func (p *dnsCache) searchDnsCache(ip string) string {
	pCrawl := p.root
	for _, v := range []rune(ip) {
		index := p.getIndexFromRune(v)
		if pCrawl.child[index] == nil {
			return ""
		}
		pCrawl = pCrawl.child[index]
	}
	if pCrawl != nil && pCrawl.isLeaf {
		return pCrawl.url
	}
	return ""
}

func newDnsCache() *dnsCache {
	return &dnsCache{root: &tireTree{
		false,
		"",
		make([]*tireTree, charCount),
	}}
}

func main() {

}
