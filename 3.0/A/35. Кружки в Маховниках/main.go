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

type CoursesInfo struct {
	vertexOut [][]int
	vertexIn  []*Ints
	counter   []int
	completed []bool
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	var v int
	_, _ = fmt.Sscanf(line, "%d", &v)

	ci := CoursesInfo{
		vertexOut: make([][]int, v+1),
		vertexIn:  make([]*Ints, v+1),

		counter:   make([]int, v+1),
		completed: make([]bool, v+1),
	}

	for i := range ci.vertexOut {
		ci.vertexOut[i] = make([]int, 0)
	}

	for i := 1; i <= v; i++ {
		line, _ = reader.ReadString('\n')

		nums := strings.Fields(line)
		if nums[0] != "0" {
			ci.counter[i], _ = strconv.Atoi(nums[0])
			ci.vertexIn[i] = NewInts(ci.counter[i])
			heap.Init(ci.vertexIn[i])

			for j := 1; j < len(nums); j++ {
				vertex, _ := strconv.Atoi(nums[j])
				ci.vertexOut[vertex] = append(ci.vertexOut[vertex], i)
				heap.Push(ci.vertexIn[i], vertex)
			}
		}
	}

	writer := bufio.NewWriter(os.Stdout)

	for i := 1; i <= v; i++ {
		TakeCourse(i, ci, writer)
	}

	_ = writer.Flush()
}

func TakeNecessaryCourses(course int, ci CoursesInfo, ints *Ints, hash map[int]struct{}) {
	if ci.completed[course] || ci.vertexIn[course] == nil {
		return
	}

	for i := range ci.vertexIn[course].arr {
		curCourse := ci.vertexIn[course].arr[i]

		if _, ok := hash[curCourse]; !ok {
			hash[curCourse] = struct{}{}

			heap.Push(ints, curCourse)

			TakeNecessaryCourses(curCourse, ci, ints, hash)
		}
	}
}

func TakeCourse(course int, ci CoursesInfo, writer *bufio.Writer) {
	if ci.completed[course] {
		return
	}

	ints := NewInts(0)
	heap.Init(ints)
	TakeNecessaryCourses(course, ci, ints, make(map[int]struct{}))

	for ints.Len() > 0 {
		vertex := heap.Pop(ints).(int)
		TakeCourse(vertex, ci, writer)
	}

	for ci.counter[course] != 0 {
		vertex := heap.Pop(ci.vertexIn[course]).(int)
		TakeCourse(vertex, ci, writer)
	}

	ci.completed[course] = true

	for _, courseIndex := range ci.vertexOut[course] {
		ci.counter[courseIndex]--
	}

	_, _ = fmt.Fprintf(writer, "%d ", course)
	return
}
