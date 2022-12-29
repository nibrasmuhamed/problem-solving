package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Queue struct {
	front *Node
	rear  *Node
}

func (q *Queue) enqueue(data int) {
	newNode := &Node{data, nil}
	// rep := new(Node)
	// rep.data = 19
	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
		return
	}
	q.rear.next = newNode
	q.rear = newNode
}

func main() {
	myQueue := Queue{}
	myQueue.enqueue(8)
	myQueue.enqueue(7)
	myQueue.enqueue(5)
	myQueue.enqueue(3)
	myQueue.dequeue()
	myQueue.Print()
}

func (q *Queue) dequeue() {
	if q.front == nil {
		fmt.Println("queue empty")
		return
	}
	q.front = q.front.next
}

func (q Queue) Print() {
	if q.rear == nil {
		fmt.Println("queue is empty")
		return
	}
	temp := q.front
	for temp != nil {
		fmt.Printf("%d - ", temp.data)
		temp = temp.next
	}
}
