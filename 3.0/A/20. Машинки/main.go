package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CarCount struct {
	next int
	car  int
}

type CarCounter struct {
	arr  []CarCount
	hash map[int]int
}

func NewCarCounter(capacity int) *CarCounter {
	return &CarCounter{
		arr:  make([]CarCount, 0, capacity),
		hash: make(map[int]int),
	}
}

func (s *CarCounter) Len() int {
	return len(s.arr)
}

func (h *CarCounter) Push(element CarCount) {
	h.arr = append(h.arr, element)
	pos := len(h.arr) - 1

	h.hash[element.car] = pos
	h.shiftUp(pos)
}

func (h *CarCounter) shiftUp(pos int) int {
	prev := (pos - 1) / 2

	for pos > 0 && h.arr[pos].next > h.arr[prev].next {
		h.Swap(pos, prev)

		pos = prev
		prev = (pos - 1) / 2
	}

	return pos
}

func (h *CarCounter) shiftDown(pos int) int {
	for pos*2+1 < len(h.arr) {
		next := -1

		if h.arr[pos*2+1].next > h.arr[pos].next {
			next = pos*2 + 1
		}

		if pos*2+2 < len(h.arr) && h.arr[pos*2+2].next > h.arr[pos].next && h.arr[pos*2+2].next > h.arr[pos*2+1].next {
			next = pos*2 + 2
		}

		if next == -1 {
			return pos
		}

		h.Swap(next, pos)
		pos = next
	}

	return pos
}

func (h *CarCounter) Swap(i1, i2 int) {
	h.arr[i1], h.arr[i2] = h.arr[i2], h.arr[i1]

	if h.hash[h.arr[i1].car] == i2 {
		h.hash[h.arr[i1].car] = i1
	}

	if h.hash[h.arr[i2].car] == i1 {
		h.hash[h.arr[i2].car] = i2
	}
}

func (h *CarCounter) Change(carCount CarCount) {
	h.arr[h.hash[carCount.car]].next = carCount.next

	index := h.shiftUp(h.hash[carCount.car])
	index = h.shiftDown(index)

	h.hash[carCount.car] = index
}

func (h *CarCounter) Delete(index int) {
	delete(h.hash, h.arr[index].car)

	h.arr[index] = h.arr[len(h.arr)-1]
	h.hash[h.arr[index].car] = index
	h.arr = h.arr[:len(h.arr)-1]

	index = h.shiftUp(index)
	index = h.shiftDown(index)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	var n, k, p int
	_, _ = fmt.Sscanf(line, "%d %d %d", &n, &k, &p)

	cars := make([]int, p)

	for i := 0; i < p; i++ {
		line, _ = reader.ReadString('\n')
		cars[i], _ = strconv.Atoi(strings.TrimRight(line, "\r\n"))
	}

	hashCar := make(map[int]int)
	nextCar := make([]int, p)

	for i := p - 1; i >= 0; i-- {
		if index, ok := hashCar[cars[i]]; ok {
			nextCar[i] = index
		} else {
			nextCar[i] = p
		}

		hashCar[cars[i]] = i
	}

	for i := range hashCar {
		delete(hashCar, i)
	}

	carsCounter := NewCarCounter(k)

	actions := 0

	for i := 0; i < p; i++ {
		if _, ok := carsCounter.hash[cars[i]]; ok {
			carsCounter.Change(CarCount{next: nextCar[i], car: cars[i]})
		} else if len(carsCounter.arr) < k {
			carsCounter.Push(CarCount{next: nextCar[i], car: cars[i]})
			actions++
		} else {
			carsCounter.Delete(0)
			carsCounter.Push(CarCount{next: nextCar[i], car: cars[i]})
			actions++
		}
	}

	fmt.Println(actions)
}
