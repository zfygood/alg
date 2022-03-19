package main

import (
	"sync"
)

//数组查找元素
func findDupByXOR(arr []int) int {
	if arr == nil {
		return -1
	}

	r := 0
	l := len(arr)

	for _, v := range arr {
		r ^= v
	}

	for i := 0; i < l; i++ {
		r ^= i
	}
	return r
}

func findDupByMap(arr []int) int {
	if arr == nil {
		return -1
	}
	l := len(arr)
	i := 0
	index := 0
	for {
		if arr[i] >= l {
			return -1
		}
		if arr[index] < 0 {
			break
		}
		arr[index] *= -1
		index = arr[index] * -1
		if index >= l {
			return -1
		}
	}
	return index
}

type set struct {
	m map[interface{}]bool
	sync.RWMutex
}

func (s *set) add(item interface{}) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func (s *set) remove(item interface{}) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

func (s *set) contains(item interface{}) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}
func (s *set) isEmpty() bool {
	return s.len() == 0
}

func (s *set) list() []interface{} {
	s.RLock()
	defer s.RUnlock()
	var list []interface{}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
func (s *set) len() int {
	return len(s.list())
}

func findDup(arr []int, num int) *set {
	s := &set{
		m:       map[interface{}]bool{},
		RWMutex: sync.RWMutex{},
	}
	if arr == nil {
		return s
	}

	l := len(arr)
	index := arr[0]
	num = num - 1
	for {
		if arr[index] < 0 {
			num--
			arr[index] = l - num
			s.add(index)
		}
		if num == 0 {
			return s
		}
		arr[index] *= -1
		index = arr[index] * -1
	}

}

func main() {

}
