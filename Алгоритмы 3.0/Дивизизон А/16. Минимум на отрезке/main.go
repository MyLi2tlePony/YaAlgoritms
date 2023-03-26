package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Heap struct {
	arr  []int
	hash map[int][]int
}

func NewHeap() *Heap {
	return &Heap{
		arr:  make([]int, 0),
		hash: make(map[int][]int),
	}
}

func (h *Heap) Push(element int) {
	h.arr = append(h.arr, element)
	pos := len(h.arr) - 1

	h.hash[element] = append(h.hash[element], pos)
	h.shiftUp(pos)
}

func (h *Heap) shiftUp(pos int) int {
	prev := (pos - 1) / 2

	for pos > 0 && h.arr[pos] < h.arr[prev] {
		h.Swap(pos, prev)

		pos = prev
		prev = (pos - 1) / 2
	}

	return pos
}

func (h *Heap) shiftDown(pos int) int {
	for pos*2+1 < len(h.arr) {
		next := -1

		if h.arr[pos*2+1] < h.arr[pos] {
			next = pos*2 + 1
		}

		if pos*2+2 < len(h.arr) && h.arr[pos*2+2] < h.arr[pos] && h.arr[pos*2+2] < h.arr[pos*2+1] {
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

func (h *Heap) Swap(i1, i2 int) {
	h.arr[i1], h.arr[i2] = h.arr[i2], h.arr[i1]

	for i := range h.hash[h.arr[i1]] {
		if h.hash[h.arr[i1]][i] == i2 {
			h.hash[h.arr[i1]][i] = i1
			break
		}
	}

	for i := range h.hash[h.arr[i2]] {
		if h.hash[h.arr[i2]][i] == i1 {
			h.hash[h.arr[i2]][i] = i2
			break
		}
	}
}

func (h *Heap) Change(old, new int) {
	index := h.hash[old][len(h.hash[old])-1]
	h.hash[old] = h.hash[old][:len(h.hash[old])-1]

	h.arr[index] = new

	index = h.shiftUp(index)
	index = h.shiftDown(index)

	h.hash[new] = append(h.hash[new], index)
}

func (h *Heap) First() int {
	if len(h.arr) > 0 {
		return h.arr[0]
	}

	return -1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	var n, k int
	_, _ = fmt.Sscanf(line, "%d %d\r\n", &n, &k)

	nums := make([]int, n)
	line, _ = reader.ReadString('\n')

	for i, number := range strings.Fields(line) {
		num, _ := strconv.Atoi(number)
		nums[i] = num
	}

	window := NewHeap()
	for i := 0; i < k; i++ {
		window.Push(nums[i])
	}

	writer := bufio.NewWriter(os.Stdout)
	_, _ = fmt.Fprintln(writer, window.First())

	for i := k; i < len(nums); i++ {
		window.Change(nums[i-k], nums[i])
		_, _ = fmt.Fprintln(writer, window.First())
	}

	_ = writer.Flush()
}
