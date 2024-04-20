package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(numberOfMinimunCandies([]int{705}, 895))
}
func numberOfMinimunCandies(A []int, B int) int {
	h := &MinHeapInt{}
	for _, v := range A {
		heap.Push(h, v)
	}

	eatenCandies := 0

	for h.Len() > 0 {
		candies := heap.Pop(h).(int)

		if candies > B {
			heap.Push(h, candies) // Put remaining candies back in the heap
			break
		}

		eatenCandies += candies / 2

		if h.Len() > 0 {
			remainingCandies := heap.Pop(h).(int)
			heap.Push(h, candies+remainingCandies)
		}
	}

	return eatenCandies
}

type MinHeapInt []int

func (h MinHeapInt) Len() int           { return len(h) }
func (h MinHeapInt) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeapInt) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeapInt) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func solve2(A []int, B int) int {
	pq := &Heap{}
	heap.Init(pq)
	//Let's keep a state array to keep track of the value of every element in the array after K operations.
	//Initially state array will be the same as the inital array.
	state := make([]int, len(A))
	copy(state, A)
	//Consider a min heap. And initially push the next state of every element in the heap.
	//Note that you need to keep track of the indices of every element in the heap, present in the initial array.
	for i, _ := range A {
		heap.Push(pq, pair{2 * A[i], i})
	}
	for B > 0 {
		B--
		top := heap.Pop(pq).(pair)
		//Pick the top element, and change the state of that element, in the state array.
		//Pop this element and push the next state in the heap.
		state[top.y] = top.x
		heap.Push(pq, pair{A[top.y] + top.x, top.y})
	}
	mx := state[0]
	for i, _ := range state {
		if state[i] > mx {
			mx = state[i]
		}
	}
	return mx
}
func solve(A []int, B int) int {
	h := &Heap{}
	heap.Init(h)
	for i := 0; i < len(A); i++ {
		x := pair{y: i, x: A[i] * 2}
		heap.Push(h, x)
	}
	dup := make([]int, len(A))
	copy(dup, A)
	for i := 0; i < B; i++ {
		top := heap.Pop(h).(pair)
		dup[top.y] = top.x
		// top.x = top.x + A[top.y]

		// dup[x.Index] += A[x.Index]
		heap.Push(h, pair{x: top.x + A[top.y], y: top.y})
	}
	min := 1<<31 - 1
	for i := 0; i < len(A); i++ {
		x := dup[i]
		if x < min {
			min = x
		}
	}
	return min
}

type MyStruct struct {
	Index int
	Value int
}

// MinHeap is a min-heap implementation for MyStruct
type MinHeap []MyStruct

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return h[i].Value < h[j].Value
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(MyStruct))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type pair struct {
	x, y int
}
type Heap []pair

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].x < h[j].x }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(pair))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/**
 * @input A : Integer array
 * @input B : Integer
 *
 * @Output Integer
 */
