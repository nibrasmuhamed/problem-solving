package main

import (
	"fmt"
)

type Queue struct {
	items []string
}

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(item string) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the element from the front of the queue
func (q *Queue) Dequeue() string {
	if len(q.items) == 0 {
		// fmt.Prstringln("Queue is empty")
		return "" // You can choose a different approach for empty queue, like returning an error
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// Peek returns the element from the front of the queue without removing it
func (q *Queue) Peek() string {
	if len(q.items) == 0 {
		// fmt.Prstringln("Queue is empty")
		return "" // You can choose a different approach for empty queue, like returning an error
	}

	return q.items[0]
}
func (q *Queue) EnqueueFront(item string) {
	newItems := make([]string, len(q.items)+1)
	newItems[0] = item
	copy(newItems[1:], q.items)
	q.items = newItems
}

// IsEmpty returns true if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
func fistNonRepeating(A string) string {
	q := Queue{}
	m := map[byte]int{}
	s := ""
	for i := 0; i < len(A); i++ {
		_, ok := m[A[i]]
		if !ok {
			m[A[i]] = 1
		} else {
			m[A[i]]++
		}
		for !q.IsEmpty() && m[byte(q.Peek()[0])] > 1 {
			q.Dequeue()
		}
		if m[A[i]] == 1 {
			q.Enqueue(string(A[i]))
			if q.Peek() == "#" {
				q.Dequeue()
			}
		} else if m[A[i]] > 1 && q.IsEmpty() {
			q.Enqueue("#")
		}
		s += q.Peek()
	}
	return s
}

func main() {
	fmt.Println(fistNonRepeating("xxikrwmjvsvckfrqxnibkcasompsuyuogauacjrr"))
}

// func reverse(A []int, B int) []int {
// 	var res []int
// 	q := Queue{}
// 	for i := 0; i < len(A); i++ {
// 		q.items = append(q.items, A[i])
// 	}

// 	for i := 0; i < B; i++ {
// 		res = append(res, q.Dequeue())
// 	}
// 	for _, v := range res {
// 		q.EnqueueFront(v)
// 	}
// 	return q.items
// }
