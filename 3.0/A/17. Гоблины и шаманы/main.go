package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type listElement struct {
	next *listElement

	value int
}

type list struct {
	head *listElement
	tail *listElement
	mid  *listElement

	len int
}

func (l *list) Pop() int {
	value := l.head.value
	l.head = l.head.next

	if l.len%2 == 0 {
		l.mid = l.mid.next
	}

	l.len--
	return value
}

func (l *list) Push(value int) {
	element := listElement{
		value: value,
	}

	if l.len == 0 {
		l.head = &element
		l.mid = &element
	} else if l.len == 1 {
		l.head.next = &element
	} else {
		l.tail.next = &element

		if l.len%2 == 0 {
			l.mid = l.mid.next
		}
	}

	l.tail = &element

	l.len++
}

func (l *list) PushMid(value int) {
	element := listElement{
		value: value,
	}

	if l.len == 0 {
		l.head = &element
		l.mid = &element
		l.tail = &element
	} else if l.len == 1 {
		l.tail = &element
		l.head.next = l.tail
	} else {
		element.next = l.mid.next
		l.mid.next = &element

		if l.len%2 == 0 {
			l.mid = l.mid.next
		}
	}

	l.len++
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimRight(line, "\r\n"))

	queue := list{}

	for i := 0; i < n; i++ {
		line, _ = reader.ReadString('\n')

		if strings.HasPrefix(line, "-") {
			fmt.Println(queue.Pop())
		} else if strings.HasPrefix(line, "*") {
			var value int
			_, _ = fmt.Sscanf(line, "* %d", &value)

			queue.PushMid(value)
		} else if strings.HasPrefix(line, "+") {
			var value int
			_, _ = fmt.Sscanf(line, "+ %d", &value)

			queue.Push(value)
		}
	}
}
