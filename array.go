package main

import (
	"fmt"
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

func getMaxAndMin(arr []int) (max, min int) {
	if arr == nil {
		return
	}
	l := len(arr)
	max = arr[0]
	min = arr[0]

	for i := 0; i < l-1; i = i + 2 {
		if arr[i] > arr[i+1] {
			tmp := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = tmp
		}
	}
	for i := 0; i < l; i = i + 2 {
		if arr[i] < min {
			min = arr[i]
		}
	}

	for i := 1; i < l; i = i + 2 {
		if arr[i] > max {
			max = arr[i]
		}
	}

	if l%2 == 1 {
		if max < arr[l-1] {
			max = arr[l-1]
		}
		if min > arr[l-1] {
			min = arr[l-1]
		}
	}
	return
}

func getMaxAndMinRe(arr []int, l, r int) (max, min int) {
	if arr == nil {
		return
	}
	m := (l + r) / 2

	if l == r {
		max = arr[l]
		min = arr[l]
		return
	}

	if l+1 == r {
		if arr[l] >= arr[r] {
			max = arr[l]
			min = arr[r]
		} else {
			max = arr[r]
			min = arr[l]
		}
		return
	}
	lMax, lMin := getMaxAndMinRe(arr, l, m)
	rMax, rMin := getMaxAndMinRe(arr, m+1, r)
	if lMax > rMax {
		max = lMax
	} else {
		max = rMax
	}

	if lMin > rMin {
		min = rMin
	} else {
		min = lMin
	}
	return
}
func getMinPara(arr []int, low, high int) int {
	if high < low {
		return arr[0]
	}
	if high == low {
		return arr[low]
	}

	mid := low + ((high - low) >> 1)
	if arr[mid] < arr[mid-1] {
		return arr[mid]
	} else if arr[mid+1] < arr[mid] {
		return arr[mid+1]
	} else if arr[high] > arr[mid] {
		return getMinPara(arr, low, mid-1)
	} else if arr[mid] > arr[low] {
		return getMinPara(arr, mid+1, high)
	} else {
		left := getMinPara(arr, low, mid-1)
		right := getMinPara(arr, mid+1, high)
		if left < right {
			return left
		} else {
			return right
		}
	}
}

func getMin(arr []int) int {
	if arr == nil {
		return -1
	} else {
		return getMinPara(arr, 0, len(arr)-1)
	}
}

func rotateArr(arr []int, div int) {
	if arr == nil || div <= 0 || div >= len(arr)-1 {
		return
	}
	swap(arr, 0, div)
	swap(arr, div+1, len(arr)-1)
	swap(arr, 0, len(arr)-1)

}
func swap(arr []int, low, high int) {
	for ; low < high; low, high = low+1, high-1 {
		tmp := arr[low]
		arr[low] = arr[high]
		arr[high] = tmp
	}
}

func getNum(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	a := 0
	b := 0
	for i, v := range arr {
		fmt.Println(i, v)
		a += v
		b += i
	}

	b = b + len(arr)*2 + 1
	return b - a
}

func getNumXOR(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}

	a := arr[0]
	b := 1
	for i := 1; i < len(arr); i++ {
		a ^= arr[i]
	}
	for j := 2; j <= len(arr)+1; j++ {
		b ^= j
	}
	return a ^ b

}

func main() {

}
