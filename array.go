package main

import (
	"fmt"
	"math"
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

func getNumXor(arr []int) {
	if arr == nil || len(arr) < 1 {
		return
	}
	result := 0
	position := uint(0)

	for _, v := range arr {
		result ^= v
	}
	tmpResult := result

	for i := result; i&1 == 0; i = i >> 1 {
		position++
	}

	for _, v := range arr {
		if (v>>position)&1 == 1 {
			result ^= v
		}
	}

	fmt.Println(result)

	fmt.Println(result ^ tmpResult)
}

func getNumOdd(arr []int) {
	if arr == nil || len(arr) < 1 {
		return
	}
	data := map[int]int{}
	for _, v := range arr {
		if r, ok := data[v]; ok {
			if r == 1 {
				data[v] = 0
			} else {
				data[v] = 1
			}
		} else {
			data[v] = 1
		}
	}

	for _, v := range arr {
		if data[v] == 1 {
			fmt.Println(v)
		}
	}

}

func findSmallK(arr []int, l, h, k int) int {
	i := l
	j := h
	tmp := arr[i]
	for i < j {
		for i < j && arr[j] >= tmp {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}
		for i < j && arr[i] <= tmp {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = tmp
	index := i - l
	if index == k-1 {
		return tmp
	} else if index > k-1 {
		return findSmallK(arr, l, i-1, k)
	} else {
		return findSmallK(arr, i+1, h, k-(i-l)-1)
	}

}

func findSmallKK(arr []int, k, j int) int {
	j++
	tmp := arr[0]
	//fmt.Println(arr)
	for i := 1; i < len(arr); i++ {
		if tmp > arr[i] {
			arr[0] = arr[i]
			arr[i] = tmp
			tmp = arr[0]
		}
	}
	if j == k {
		return arr[0]
	}
	return findSmallKK(arr[1:], k, j)
}

func findTop(arr []int) (r1, r2, r3 int) {
	if arr == nil || len(arr) < 3 {
		return
	}

	r1 = math.MinInt64
	r2 = math.MinInt64
	r3 = math.MinInt64

	for _, v := range arr {
		if v > r1 {
			r3 = r2
			r2 = r1
			r1 = v
			fmt.Println(r1, r2, r3)
		} else if v > r2 && v != r1 {
			r3 = r2
			r2 = v
			fmt.Println(r1, r2, r3)
		} else if v > r3 && v != r2 {
			r3 = v
			fmt.Println(r1, r2, r3)
		}
	}
	return
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minSan(a, b, c int) int {
	var m int
	if a < b {
		m = a
	} else {
		m = b
	}
	if m > c {
		m = c
	}
	return m
}
func max(a, b, c int) (m int) {
	if a < b {
		m = b
	} else {
		m = a
	}
	if m < c {
		m = c
	}
	return
}

func abs(m int) int {
	if m < 0 {
		return -m
	}
	return m
}

func minDistance(arr []int, num1, num2 int) int {
	if arr == nil || len(arr) <= 0 {
		return math.MaxInt64
	}

	lastPos1 := -1
	lastPos2 := -1
	minDis := math.MaxInt64
	for i, v := range arr {
		if v == num1 {
			lastPos1 = i
			if lastPos2 >= 0 {
				minDis = min(minDis, lastPos1-lastPos2)
			}
		}

		if v == num2 {
			lastPos2 = i
			if lastPos1 >= 0 {
				minDis = min(minDis, lastPos2-lastPos1)
			}
		}
	}
	return minDis
}

func minDistanceSan(a, b, c []int) int {
	aLen := len(a)
	bLen := len(b)
	cLen := len(c)
	minDist := math.MaxInt64
	i, j, k := 0, 0, 0
	for {
		curDist := max(abs(a[i]-b[j]), abs(b[j]-c[k]), abs(a[i]-c[k]))
		if curDist < minDist {
			minDist = curDist
		}
		m := minSan(a[i], b[j], c[k])
		if m == a[i] {
			i++
			if i >= aLen {
				break
			}
		} else if m == b[j] {
			j++
			if j >= bLen {
				break
			}
		} else {
			k++
			if k >= cLen {
				break
			}
		}
	}
	return minDist
}

func findMinMath(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	begin := 0
	mid := 0
	end := len(arr) - 1
	var absMin int
	if arr[0] > 0 {
		return arr[0]
	}
	if arr[end] <= 0 {
		return arr[end]
	}

	for {
		mid = begin + ((end - begin) >> 1)
		if arr[mid] == 0 {
			return 0
		} else if arr[mid] > 0 {
			if arr[mid-1] > 0 {
				end = mid - 1
			} else if arr[mid-1] == 0 {
				return 0
			} else {
				break
			}
		} else {
			if arr[mid+1] < 0 {
				begin = mid + 1
			} else if arr[mid+1] == 0 {
				return 0
			} else {
				break
			}
		}
	}

	if arr[mid] > 0 {
		if arr[mid] < abs(arr[mid-1]) {
			absMin = arr[mid]
		} else {
			absMin = arr[mid-1]
		}
	} else {
		if abs(arr[mid]) < arr[mid+1] {
			absMin = arr[mid]
		} else {
			absMin = arr[mid+1]
		}
	}
	return absMin
}

func findMin(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}

	absMin := math.MaxInt64
	for _, v := range arr {
		if abs(v) < abs(absMin) {
			absMin = v
		}
	}
	return absMin
}

func maxSubArraySum(arr []int) int {
	if arr == nil || len(arr) < 1 {
		return -1
	}
	l := len(arr)
	maxSum := math.MinInt64
	for i := 0; i < l; i++ {
		sum := 0
		for j := i; j < l; j++ {
			sum += arr[j]
			if sum > maxSum {
				maxSum = sum
			}
			fmt.Println(arr[j], sum, maxSum)

		}
	}
	return maxSum
}
func MaxRe(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxSubArraySumD(arr []int) int {
	nAll := arr[0]
	nEnd := arr[0]
	for _, v := range arr {
		fmt.Println(nEnd+v, v)
		nEnd = MaxRe(nEnd+v, v)
		fmt.Println(nEnd)
		nAll = MaxRe(nEnd, nAll)
		fmt.Println(nAll, nEnd)
	}
	return nAll
}

func maxSubArrayEx(arr []int) (maxSum, begin, end int) {
	maxSum = math.MinInt64
	nSum := 0
	nStart := 0

	for i, v := range arr {
		if nSum < 0 {
			nSum = v
			nStart = i
		} else {
			nSum += v
		}

		if nSum > maxSum {
			maxSum = nSum
			begin = nStart
			end = i
		}
	}
	return
}

func main() {

}
