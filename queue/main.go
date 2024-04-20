package main

import (
	"container/list"
	"fmt"
	"strconv"
)

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

type Queue struct {
	Front *Node
	Rear  *Node
}

func (d *Queue) IsEmpty() bool {
	return d.Front == nil && d.Rear == nil
}

func (d *Queue) EnqueueFront(val int) {
	newNode := &Node{Val: val}
	if d.IsEmpty() {
		d.Front = newNode
		d.Rear = newNode
	} else {
		newNode.Next = d.Front
		d.Front.Prev = newNode
		d.Front = newNode
	}
}

func (d *Queue) EnqueueRear(val int) {
	newNode := &Node{Val: val}
	if d.IsEmpty() {
		d.Front = newNode
		d.Rear = newNode
	} else {
		newNode.Prev = d.Rear
		d.Rear.Next = newNode
		d.Rear = newNode
	}
}

func (d *Queue) DequeueFront() int {
	if d.IsEmpty() {
		return -1
	}

	ret := d.Front.Val
	d.Front = d.Front.Next

	if d.Front == nil {
		d.Rear = nil
	} else {
		d.Front.Prev = nil
	}

	return ret
}

func (d *Queue) DequeueRear() int {
	if d.IsEmpty() {
		return -1
	}

	ret := d.Rear.Val
	d.Rear = d.Rear.Prev

	if d.Rear == nil {
		d.Front = nil
	} else {
		d.Rear.Next = nil
	}

	return ret
}
func (d *Queue) PeekFront() int {
	if d.IsEmpty() {
		return -1
	}
	return d.Front.Val
}

func (d *Queue) PeekRear() int {
	if d.IsEmpty() {
		return -1
	}
	return d.Rear.Val
}
func allNumbers(A int) []int {
	q := Queue{}
	res := []int{}

	q.EnqueueRear(1)
	q.EnqueueRear(2)
	q.EnqueueRear(3)
	for i := 1; i < A; i++ {
		x := q.DequeueFront()
		res = append(res, x)
		q.EnqueueRear(x*10 + 1)
		q.EnqueueRear(x*10 + 2)
		q.EnqueueRear(x*10 + 3)
	}
	res = append(res, q.DequeueFront())
	return res
}

// func main() {
// 	fmt.Println(allNumbers(4))
// 	// q := Queue{}
// 	// fmt.Println("Is empty: ", q.IsEmpty()) // true
// 	// q.Enqueue(2)
// 	// q.Enqueue(3)
// 	// q.E

//		// q.Print()s
//	}
func sumOfAllSubArray(A []int, B int) int {
	mod := 10000000007
	maxElement := Queue{}
	minElement := Queue{}

	out := 0

	for i := 0; i < B; i++ {
		for !maxElement.IsEmpty() && A[maxElement.PeekRear()] < A[i] {
			maxElement.DequeueRear()
		}
		for !minElement.IsEmpty() && A[minElement.PeekRear()] > A[i] {
			minElement.DequeueRear()
		}
		minElement.EnqueueRear(i)
		maxElement.EnqueueRear(i)
	}
	// fmt.Println(minElement.PeekFront())
	for i := B; i < len(A); i++ {
		out = out + (((A[maxElement.PeekFront()] % mod) + (A[minElement.PeekFront()] % mod)) % mod)

		for !maxElement.IsEmpty() && maxElement.PeekFront() <= i-B {
			maxElement.DequeueFront()
		}
		for !minElement.IsEmpty() && minElement.PeekFront() <= i-B {
			minElement.DequeueFront()
		}
		for !maxElement.IsEmpty() && A[maxElement.PeekRear()] < A[i] {
			maxElement.DequeueRear()
		}
		for !minElement.IsEmpty() && A[minElement.PeekRear()] > A[i] {
			minElement.DequeueRear()
		}
		maxElement.EnqueueRear(i)
		minElement.EnqueueRear(i)
	}

	out = out + (((A[maxElement.PeekFront()] % mod) + (A[minElement.PeekFront()] % mod)) % mod)
	return out
}

func slidingMaximum(A []int, B int) []int {
	maxElement := Queue{}
	out := []int{}

	for i := 0; i < B; i++ {
		for !maxElement.IsEmpty() && A[maxElement.PeekRear()] <= A[i] {
			maxElement.DequeueRear()
		}
		maxElement.EnqueueRear(i)
	}

	for i := B; i < len(A); i++ {
		out = append(out, A[maxElement.PeekFront()])

		for !maxElement.IsEmpty() && maxElement.PeekFront() <= i-B {
			maxElement.DequeueFront()
		}

		for !maxElement.IsEmpty() && A[maxElement.PeekRear()] <= A[i] {
			maxElement.DequeueRear()
		}
		maxElement.EnqueueRear(i)
	}

	out = append(out, A[maxElement.PeekFront()])
	return out
}

func findPerfectNumber(A int) string {
	q := Queue{}
	// if A == 1 {
	// 	return "11"
	// } else if A == 2 {
	// 	return "22"
	// }
	q.EnqueueRear(1)
	q.EnqueueRear(2)
	for i := 1; i < A; i++ {
		x := q.DequeueFront()
		u := x*10 + 1
		v := x*10 + 2
		q.EnqueueRear(u)
		q.EnqueueRear(v)
	}
	j := strconv.Itoa(q.DequeueFront())
	return j + reverseString(j)
}
func reverseString(input string) string {
	// Convert the string to a slice of runes
	runes := []rune(input)

	// Get the length of the slice
	length := len(runes)

	// Swap the characters from the beginning and end until reaching the middle
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the slice back to a string
	return string(runes)
}

func solve(A []int, B []int) int {
	q := list.New()
	for _, v := range A {
		q.PushBack(v)
	}
	i, count := 0, 0
	for q.Len() != 0 {
		node := q.Front()
		q.Remove(node)
		if node.Value.(int) == B[i] {
			i++
		} else {
			q.PushBack(node.Value.(int))
		}
		count++
	}
	return count
}

func main() {
	fmt.Println(solve([]int{2, 3, 1, 5, 4}, []int{1, 3, 5, 4, 2}))
}
