package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ints struct {
	arr []int
}

func NewInts(capacity int) *Ints {
	return &Ints{
		arr: make([]int, 0, capacity),
	}
}

func (s *Ints) Len() int {
	return len(s.arr)
}

func (s *Ints) Less(i, j int) bool {
	return s.arr[i] < s.arr[j]
}

func (s *Ints) Swap(i, j int) {
	s.arr[i], s.arr[j] = s.arr[j], s.arr[i]
}

func (s *Ints) Push(x any) {
	s.arr = append(s.arr, x.(int))
}

func (s *Ints) Pop() any {
	item := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return item
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	var n int
	_, _ = fmt.Sscanf(line, "%d", &n)

	nums := NewInts(n)
	heap.Init(nums)

	line, _ = reader.ReadString('\n')
	for _, num := range strings.Fields(line) {
		number, _ := strconv.Atoi(num)
		heap.Push(nums, number)
	}

	sum := 0

	for nums.Len() > 1 {
		num1 := heap.Pop(nums).(int)
		num2 := heap.Pop(nums).(int)
		sum += num1 + num2

		heap.Push(nums, num1+num2)
	}

	sum = sum * 5

	fmt.Printf("%d.%02d", sum/100, sum%100)
}

func sortLast(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			nums[i], nums[i-1] = nums[i-1], nums[i]
		}
	}
}
