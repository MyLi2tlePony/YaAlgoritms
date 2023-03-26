package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Schedule struct {
	parkingIndex  int
	departureTime int
}

type Schedules struct {
	arr []Schedule
}

func NewSchedules(capacity int) *Schedules {
	return &Schedules{
		arr: make([]Schedule, 0, capacity),
	}
}

func (s *Schedules) Len() int {
	return len(s.arr)
}

func (s *Schedules) Less(i, j int) bool {
	return s.arr[i].departureTime < s.arr[j].departureTime
}

func (s *Schedules) Swap(i, j int) {
	s.arr[i], s.arr[j] = s.arr[j], s.arr[i]
}

func (s *Schedules) Push(x any) {
	s.arr = append(s.arr, x.(Schedule))
}

func (s *Schedules) Pop() any {
	item := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return item
}

func (s *Schedules) First() (Schedule, bool) {
	if s.Len() < 1 {
		return Schedule{}, false
	}

	return s.arr[0], true
}

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

func (s *Ints) First() (int, bool) {
	if s.Len() < 1 {
		return 0, false
	}

	return s.arr[0], true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	var k, n int
	_, _ = fmt.Sscanf(line, "%d %d", &k, &n)

	parking := NewInts(k)
	for i := 0; i < k; i++ {
		parking.Push(i)
	}

	schedules := NewSchedules(n)

	heap.Init(parking)
	heap.Init(schedules)

	history := make([]int, n)

	for i := 0; i < n; i++ {
		line, _ = reader.ReadString('\n')

		var start, end int
		_, _ = fmt.Sscanf(line, "%d %d", &start, &end)

		for schedules.Len() > 0 {
			if schedule, _ := schedules.First(); schedule.departureTime < start {
				heap.Pop(schedules)
				heap.Push(parking, schedule.parkingIndex)
			} else {
				break
			}
		}

		if parking.Len() > 0 {
			freeParking := heap.Pop(parking).(int)
			heap.Push(schedules, Schedule{departureTime: end, parkingIndex: freeParking})
			history[i] = freeParking + 1
		} else {
			fmt.Print(0, i+1)
			return
		}
	}

	writer := bufio.NewWriter(os.Stdout)

	for i := range history {
		if i < n {
			_, _ = fmt.Fprintf(writer, "%d ", history[i])
		} else {
			_, _ = fmt.Fprintf(writer, "%d", history[i])
		}
	}

	_ = writer.Flush()
}
