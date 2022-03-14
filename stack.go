package main

import (
	"errors"
	"fmt"
	"math"
	"sync"
)

//数组形式得stack
type sliceStack struct {
	arr       []int
	stackSize int
}

func (p *sliceStack) IsEmpty() bool {
	return p.stackSize == 0
}

func (p *sliceStack) size() int {
	return p.stackSize
}

func (p *sliceStack) top() int {
	if p.IsEmpty() {
		panic(errors.New(""))
	}
	return p.arr[p.stackSize-1]
}

func (p *sliceStack) pop() int {
	if p.stackSize > 0 {
		p.stackSize--
		ret := p.arr[p.stackSize]
		p.arr = p.arr[:p.stackSize]
		return ret
	}
	panic(errors.New(""))
}
func (p *sliceStack) push(t int) {
	p.arr = append(p.arr, t)
	p.stackSize = p.stackSize + 1
}

func sliceMode() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	s := &sliceStack{arr: make([]int, 0)}
	s.push(1)
	fmt.Println(s.top())
	fmt.Println(s.size())
	s.pop()
	fmt.Println(s.size())

}

type LNodeL struct {
	Data interface{}
	Next *LNodeL
}
type listStack struct {
	head *LNodeL
}

//链表实现stack
func (p *listStack) isEmpty() bool {
	return p.head.Next == nil
}

func (p *listStack) size() int {
	i := 0
	tmp := p.head.Next
	for tmp != nil {
		i++
		tmp = tmp.Next
	}
	return i
}

func (p *listStack) push(i int) {
	node := &LNodeL{Data: i, Next: p.head.Next}
	p.head.Next = node
}

func (p *listStack) pop() int {
	tmp := p.head.Next
	if tmp != nil {
		p.head.Next = tmp.Next
		return tmp.Data.(int)
	}
	panic(errors.New(""))
}

func (p *listStack) top() int {
	if p.head.Next != nil {
		return p.head.Next.Data.(int)
	}
	panic(errors.New(""))
}

func listMode() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	l := &listStack{head: &LNodeL{}}
	l.push(1)
	fmt.Println(l.top())
	fmt.Println(l.size())

	l.pop()
	fmt.Println(l.size())
}

//数组实现队列
type sliceQueue struct {
	arr   []int
	front int
	rear  int
}

func (p *sliceQueue) isEmpty() bool {
	return p.front == p.rear
}
func (p *sliceQueue) size() int {
	return p.rear - p.front
}
func (p *sliceQueue) getFront() int {
	if p.isEmpty() {
		panic(errors.New(""))
	}
	return p.arr[p.front]
}

func (p *sliceQueue) getBack() int {
	if p.isEmpty() {
		panic(errors.New(""))
	}
	return p.arr[p.rear-1]
}
func (p *sliceQueue) deQueue() {
	if p.rear > p.front {
		p.rear--
		p.arr = p.arr[1:]
	} else {
		panic(errors.New(""))
	}
}

func (p *sliceQueue) enQueue(i int) {
	p.arr = append(p.arr, i)
	p.rear++
}

func sliceModeQ() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	sliceQ := &sliceQueue{arr: make([]int, 0)}
	sliceQ.enQueue(1)
	fmt.Println(sliceQ.getFront())
	fmt.Println(sliceQ.getBack())
	fmt.Println(sliceQ.size())
}

//链表实现队列
type listQueue struct {
	head *LNodeL
	end  *LNodeL
}

func (p *listQueue) isEmpty() bool {
	return p.head == nil
}
func (p *listQueue) size() int {
	i := 0
	tmp := p.head
	for tmp != nil {
		i++
		tmp = tmp.Next
	}
	return i
}
func (p *listQueue) enQueue(i int) {
	node := &LNodeL{
		Data: i,
	}
	if p.head == nil {
		p.head = node
		p.end = node
	} else {
		p.end.Next = node
		p.end = node
	}
}

func (p *listQueue) deQueue() {
	if p.head == nil {
		panic(errors.New(""))
	}
	p.head = p.head.Next
	if p.head == nil {
		p.end = nil
	}
}

func (p *listQueue) getFront() int {
	if p.head == nil {
		panic(errors.New(""))
	}
	return p.head.Data.(int)
}
func (p *listQueue) getBack() int {
	if p.end == nil {
		panic(errors.New(""))
	}
	return p.end.Data.(int)
}
func listQueueL() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	l := &listQueue{}
	l.enQueue(1)
	fmt.Println(l.getFront())
	fmt.Println(l.getBack())
	fmt.Println(l.size())
}

type sliceStackR struct {
	arr       []interface{}
	stackSize int
	sync.RWMutex
}

func (p *sliceStackR) isEmpty() bool {
	return p.stackSize == 0
}

func (p *sliceStackR) size() int {
	return p.stackSize
}
func (p *sliceStackR) push(t interface{}) {
	p.Lock()
	defer p.Unlock()
	p.arr = append(p.arr, t)
	p.stackSize = p.stackSize + 1
}
func (p *sliceStackR) pop() interface{} {
	if p.stackSize > 0 {
		p.stackSize--
		ret := p.arr[p.stackSize]
		p.arr = p.arr[:p.stackSize]
		return ret
	}
	return nil
}
func (p *sliceStackR) top() interface{} {
	if p.stackSize > 0 {
		return p.arr[p.stackSize-1]
	}
	return nil
}

func (p *sliceStackR) list() []interface{} {
	return p.arr
}

//栈
func moveBottomToTop(s *sliceStackR) {
	if s.isEmpty() {
		return
	}
	top1 := s.pop()
	if !s.isEmpty() {
		moveBottomToTop(s)
		top2 := s.top()
		if top1.(int) > top2.(int) {
			s.pop()
			s.push(top1)
			s.push(top2)
			return
		}

	}
	s.push(top1)
}
func reverseStack(s *sliceStackR) {
	if s.isEmpty() {
		return
	}
	moveBottomToTop(s)
	top := s.pop()
	reverseStack(s)
	s.push(top)
}
func newSliceStack() *sliceStackR {
	return &sliceStackR{
		arr:       make([]interface{}, 0),
		stackSize: 0,
		RWMutex:   sync.RWMutex{},
	}
}
func createStack(list []int) *sliceStackR {
	stack := newSliceStack()
	for _, v := range list {
		stack.push(v)
	}
	return stack
}
func printStack(str string, s *sliceStackR) {
	fmt.Println(str)
	for !s.isEmpty() {
		fmt.Print(s.pop(), "")
	}
	fmt.Println()
}

func isPopSerial(push string, pop string) bool {
	lPush := len(push)
	pPush := len(pop)
	if lPush == 0 || pPush == 0 || lPush != pPush {
		return false
	}
	l := 0
	n := 0
	pushRune := []rune(push)
	popRune := []rune(pop)
	s := newSliceStack()
	for l < lPush {
		s.push(pushRune[l])
		l++
		for !s.isEmpty() && s.top() == popRune[n] {
			a := s.pop()
			fmt.Println(a)
			n++
		}
	}
	if s.isEmpty() && n == pPush {
		return true
	}

	return false
}

type MinStack struct {
	elemStack *sliceStackR
	minStack  *sliceStackR
}

func (p *MinStack) push(data int) {
	p.elemStack.push(data)

	if p.minStack.isEmpty() {
		p.minStack.push(data)
	} else {
		if data <= p.minStack.top().(int) {
			p.minStack.push(data)
		}
	}

}

func (p *MinStack) pop() int {
	topData := p.elemStack.pop().(int)
	if topData == p.min() {
		p.minStack.pop()
	}
	return topData
}

func (p *MinStack) min() int {
	if p.minStack.isEmpty() {
		return math.MaxInt32
	} else {
		return p.minStack.top().(int)
	}
}

type stackQueue struct {
	aStack *sliceStackR
	bStack *sliceStackR
}

//栈模拟队列
func (p *stackQueue) push(data int) {
	p.aStack.push(data)
}

func (p *stackQueue) pop() int {
	if p.bStack.isEmpty() {
		for !p.aStack.isEmpty() {
			p.bStack.push(p.aStack.pop())
		}
	}
	return p.bStack.pop().(int)

}

type sliceQueueR struct {
	arr []interface{}
	sync.RWMutex
}

func (p *sliceQueueR) isEmpty() bool {
	return p.size() == 0
}
func (p *sliceQueueR) size() int {
	return len(p.arr)
}
func (p *sliceQueueR) getFront() interface{} {
	if p.isEmpty() {
		return nil
	}
	return p.arr[0]
}

func (p *sliceQueueR) getBack() interface{} {
	if p.isEmpty() {
		return nil
	}
	return p.arr[p.size()-1]
}

func (p *sliceQueueR) popBack() interface{} {
	p.Lock()
	defer p.Unlock()
	if p.isEmpty() {
		return nil
	}
	ret := p.arr[p.size()-1]
	p.arr = p.arr[:p.size()-1]
	return ret
}
func (p *sliceQueueR) deQueue() interface{} {
	p.Lock()
	defer p.Unlock()
	if p.isEmpty() {
		return nil
	}
	ret := p.arr[0]
	p.arr = p.arr[1:]
	return ret
}

func (p *sliceQueueR) enQueue(item interface{}) {
	p.Lock()
	defer p.Unlock()
	p.arr = append(p.arr, item)
}

func (p *sliceQueueR) enQueueF(item interface{}) {
	p.Lock()
	defer p.Unlock()
	i := []interface{}{item}
	p.arr = append(i, p.arr[:]...)
}
func (p *sliceQueueR) remove(item interface{}) {
	p.Lock()
	defer p.Unlock()
	for k, v := range p.arr {
		if v == item {
			p.arr = append(p.arr[:k], p.arr[k+1:]...)
		}
	}
}
func (p *sliceQueueR) list() []interface{} {
	return p.arr
}

type user struct {
	id   int
	name string
	seq  int
}

type loginQueue struct {
	queue *sliceQueueR
}

func (p *loginQueue) enQueue(u *user) {
	p.queue.enQueue(u)
	u.seq = p.queue.size()
}

func (p *loginQueue) deQueue(u *user) {
	p.queue.deQueue()
	p.updateSeq()
}
func (p *loginQueue) deQueueUser(u *user) {
	p.queue.remove(u)
	p.updateSeq()
}
func (p *loginQueue) updateSeq() {
	i := 1
	for _, v := range p.queue.list() {
		u := v.(*user)
		u.seq = i
		i++
	}
}

func (p *loginQueue) printQueue() {
	for _, v := range p.queue.list() {
		fmt.Println(v.(*user))
	}
}

type hashSet struct {
	m map[interface{}]bool
	sync.RWMutex
}

func newSet() *hashSet {
	return &hashSet{m: map[interface{}]bool{}}
}
func (s *hashSet) add(item interface{}) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}
func (s *hashSet) remove(item interface{}) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}
func (s *hashSet) contains(item interface{}) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok

}
func (s *hashSet) len() int {
	return len(s.list())
}

func (s *hashSet) list() []interface{} {
	s.RLock()
	defer s.RUnlock()
	var list []interface{}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
func (s *hashSet) clear() {
	s.RLock()
	defer s.RUnlock()
	s.m = map[interface{}]bool{}
}
func (s *hashSet) isEmpty() bool {
	return len(s.list()) == 0
}

type LRU struct {
	cacheSize int
	queue     *sliceQueueR
	hastSet   *hashSet
}

func (p *LRU) isQueueFull() bool {
	return p.queue.size() == p.cacheSize
}
func (p *LRU) enQueue(item interface{}) {
	if p.isQueueFull() {
		p.hastSet.remove(p.queue.popBack())
	}
	p.queue.enQueue(item)
	p.hastSet.add(item)
}

func (p *LRU) accessPage(item int) {
	if !p.hastSet.contains(item) {
		p.enQueue(item)
	} else if p.queue.getFront() != item {
		p.queue.remove(item)
		p.queue.enQueueF(item)
	}

}

func (p *LRU) printQueue() {
	for !p.queue.isEmpty() {
		fmt.Println(p.queue.deQueue())
	}
}

//车票行程
func printResult(input map[string]string) {
	rInput := map[string]string{}
	for k, v := range input {
		rInput[v] = k
	}

	var start string
	for k, _ := range input {
		if _, ok := rInput[k]; !ok {
			start = k
		}
	}

	if start == "" {
		fmt.Println("")
	}

	to := input[start]
	//fmt.Println(start, "->", to)
	//start = to
	//to = input[to]
	for to != "" {
		fmt.Println(",", start, "->", to)
		start = to
		to = input[to]
	}
	fmt.Println()

}

type pairs struct {
	first  int
	second int
}

func findPairs(arr []int) bool {
	sumPairs := map[int]*pairs{}
	l := len(arr)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			sum := arr[i] + arr[j]
			if _, ok := sumPairs[sum]; !ok {
				sumPairs[sum] = &pairs{i, j}
			} else {
				p := sumPairs[sum]
				fmt.Print(arr[p.first], arr[p.second], arr[i], arr[j])
				fmt.Println()
				return true
			}
		}
	}
	return false
}

func main() {

}
