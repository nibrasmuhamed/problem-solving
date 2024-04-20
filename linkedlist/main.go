package main

import (
	"fmt"
	"time"
)

// Node represents a node in the linked list

// Append adds a new node with the given data to the end of the linked list
func (ll *LinkedList) Append(data int) {
	newNode := &listNode{value: data, next: nil}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	current := ll.Head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

// Display prints the elements of the linked list
func (ll *LinkedList) Display() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}
func longestPalidromeLength(A *listNode) int {
	if A == nil {
		return 0
	}
	if A.next == nil {
		return 1
	}
	count := 1
	var next, prev *listNode
	for A != nil {
		next = A.next
		A.next = prev

		count = max(count, getCount(prev, next)+1)
		count = max(count, getCount(A, next))

		prev = A
		A = next
	}
	return count
}
func getCount(a, b *listNode) int {
	count := 0
	for a != nil && b != nil {
		if a.value == b.value {
			count++
		} else {
			break
		}
		a = a.next
		b = b.next
	}
	return count
}
func printrev() {
	// loadingSymbols := []string{"|", "/", "-", "\\"}

	fmt.Print("Loading ")

	for i := 0; i < 20; i++ {
		// fmt.Print(loadingSymbols[i%len(loadingSymbols)])
		fmt.Print("...")
		time.Sleep(100 * time.Millisecond)
		fmt.Print("\v") // Move the cursor back one position
	}

	fmt.Println("\nDone!")
	// for i := 1; i <= 10; i++ {

	// 	fmt.Printf("\rCounting: -")
	// 	time.Sleep(time.Second) // Simulating some work
	// }

	// fmt.Println("\nDone!")
}

func reverseLists(A *listNode, B int) *listNode {
	count := 0
	length := 0
	temp := A
	for temp != nil {
		length++
		temp = temp.next
	}
	temp = A
	var x []*listNode
	f := A
	for i := 0; i < length; i++ {
		count++
		if count == B {
			y := temp.next
			temp.next = nil
			j := reverseList(f)
			x = append(x, j)
			temp = y
			f = y
			count = 0
			continue
		}
		temp = temp.next
	}
	newlist := &listNode{}
	temp = newlist
	for _, v := range x {
		temp.next = v
		for temp.next != nil {
			temp = temp.next
		}
	}
	return newlist.next
}
func deleteDuplicates(A *listNode) *listNode {
	temp := A
	for temp.next != nil {
		if temp.value == temp.next.value {
			temp.next = temp.next.next
			continue
		}
		temp = temp.next
	}
	return A
}

func reverseBetween(A *listNode, B int, C int) *listNode {
	if A == nil || B == C {
		return A
	}

	dummy := &listNode{value: 0, next: A}
	prev := dummy
	temp := A

	// Move to the starting point of the sublist
	for i := 1; i < B; i++ {
		prev = temp
		temp = temp.next
	}

	b := temp
	var c, d *listNode
	for i := B; i <= C; i++ {
		nextTemp := temp.next
		temp.next = d
		d = temp
		temp = nextTemp
		if i == C {
			c = temp
		}
	}

	prev.next = d
	b.next = c

	return dummy.next
}

func Reverse(A *listNode) *listNode {
	var prev, curr *listNode
	curr = A
	for curr != nil {
		nexttemp := A.next
		curr.next = prev
		prev = curr
		curr = nexttemp
	}
	return prev
}

func swapPairs(head *listNode) *listNode {
	if head == nil || head.next == nil {
		return head
	}

	dummy := &listNode{value: 0, next: head}
	prev := dummy

	for head != nil && head.next != nil {
		first := head
		second := head.next

		// Swap nodes
		prev.next = second
		first.next = second.next
		second.next = first

		// Move to the next pair
		prev = first
		head = first.next
	}

	return dummy.next
}
func swap(a int, b int) *listNode {
	c := &listNode{value: a}
	d := &listNode{value: b}
	d.next = c
	c.next = nil
	return d
}
func addTwoNumbers(A *listNode, B *listNode) *listNode {
	carry := 0
	var curr *listNode = &listNode{}
	cur := curr
	for A != nil || B != nil || carry > 0 {
		sum := carry
		if A != nil {
			sum += A.value
			A = A.next
		}
		if B != nil {
			sum += B.value
			B = B.next
		}
		carry = sum / 10
		cur.next = &listNode{value: sum % 10}
		cur = cur.next
	}

	return reverseList(curr.next)
}
func reverseList(head *listNode) *listNode {
	var prev *listNode
	current := head

	for current != nil {
		nextTemp := current.next
		current.next = prev
		prev = current
		current = nextTemp
	}

	return prev
}
func cycleNode(A *listNode) *listNode {
	var slow, fast *listNode = A, A
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			break
		}
	}
	z := A
	var pre *listNode
	for z != slow {
		pre = z
		z = z.next
		slow = slow.next
	}
	pre.next = nil
	return A
}

func sortList(A *listNode) *listNode {
	if A == nil || A.next == nil {
		return A
	}
	middle := FindMiddle(A)
	h1 := A
	h2 := middle.next
	middle.next = nil
	h1 = sortList(h1)
	h2 = sortList(h2)
	return mergeTwoLists(h1, h2)
}

func FindMiddle(A *listNode) *listNode {
	var slow, fast *listNode = A, A
	for fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	return slow
}

func mergeTwoLists(A *listNode, B *listNode) *listNode {
	var head *listNode
	if A == nil {
		return B
	}
	if B == nil {
		return A
	}
	if A.value < B.value {
		head = A
		A = A.next
	} else {
		head = B
		B = B.next
	}
	current := head
	for A != nil && B != nil {
		if A.value <= B.value {
			head.next = A
			A = A.next
			head = head.next
		} else {
			head.next = B
			B = B.next
			head = head.next
		}
	}

	if A != nil {
		head.next = A
	} else if B != nil {
		head.next = B
	}
	return current
}
func Append(head **listNode, data int) {
	newNode := &listNode{value: data, next: nil}

	if *head == nil {
		*head = newNode
		return
	}

	last := *head
	for last.next != nil {
		last = last.next
	}

	last.next = newNode

}

// Display prints the elements of the linked list
func Display(head *listNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}
func lenprint() {
	n := 0
	z := ll.Head
	for z != nil {
		n++
		z = z.next
	}
	fmt.Println("Length of the list is: ", n)
}
func revlist() {
	var prev, curr, next *listNode = nil, ll.Head, ll.Head.next
	for curr != nil {
		curr.next = prev
		prev = curr
		curr = next
		if next == nil {
			continue
		}
		next = next.next
	}
	ll.Head = prev
}

func deleteAll(n int) {
	root := ll.Head.next
	p := ll.Head
	for root != nil {
		if root.value == n {
			for root.next.value == n {
				root = root.next
			}
			p.next = root.next
			p = p.next
			root = p.next
			continue
		}
		p = p.next
		root = root.next
	}
}

type listNode struct {
	value int
	next  *listNode
}

// LinkedList represents the linked list structure
type LinkedList struct {
	Head *listNode
}

var ll LinkedList

// insert_node inserts a new node with the given value at the specified position
func insert_node(position, value int) {
	newNode := &listNode{value: value}

	if position == 1 {
		newNode.next = ll.Head
		ll.Head = newNode
		return
	}

	current := ll.Head
	for i := 1; i < position-1 && current != nil; i++ {
		current = current.next
	}

	if current == nil {
		return
	}

	newNode.next = current.next
	current.next = newNode
}

// delete_node deletes the node at the specified position
func delete_node(position int) {
	if position == 1 {
		if ll.Head == nil {
			return
		}
		ll.Head = ll.Head.next
		return
	}

	current := ll.Head
	for i := 1; i < position-1 && current != nil; i++ {
		current = current.next
	}

	if current == nil || current.next == nil {
		return
	}

	current.next = current.next.next
}

// print_ll prints the entire linked list with each element followed by a single space
func print_ll() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println()
}

func main() {
	// printrev()
	// var m1 *listNode
	// // Append(&m1, 1)
	// // Append(&m1, 1)
	// // Append(&m1, 1)
	// // Append(&m1, 1)
	// // Append(&m1, 1)
	// // Append(&m1, 1)
	// Append(&m1, 1)

	// // Append(&m1, 2)
	// // Append(&m1, 2)
	// // Append(&m1, 2)
	// // Append(&m1, 2)
	// // Append(&m1, 2)
	// // Append(&m1, 2)
	// Append(&m1, 2)

	// // Append(&m1, 3)
	// // Append(&m1, 3)
	// // Append(&m1, 3)
	// // Append(&m1, 3)
	// // Append(&m1, 3)
	// Append(&m1, 3)

	// Append(&m1, 4)
	// Append(&m1, 5)
	// Append(&m1, 6)
	// Append(&m1, 7)
	// Append(&m1, 8)
	// Append(&m1, 9)
	// Append(&m1, 10)
	// Display(reverseLists(m1, 2))
	head := listNode_new(3)
	head.right = listNode_new(4)
	head.right.down = listNode_new(7)
	head.right.down.down = listNode_new(8)

	head.right.right = listNode_new(20)
	head.right.right.down = listNode_new(22)
	head.right.right.down.down = listNode_new(28)

	head.right.right.right = listNode_new(20)
	head.right.right.right.down = listNode_new(20)
	head.right.right.right.down.down = listNode_new(39)

	head.right.right.right.right = listNode_new(30)
	head.right.right.right.right.down = listNode_new(31)
	head.right.right.right.right.down.down = listNode_new(39)
	printList(flatten(head))
}

func flatten(root *listNodex) *listNodex {
	r := new(listNodex)
	rootTracker := root

	for rootTracker != nil {
		x := rootTracker
		var z *listNodex
		if rootTracker.right != nil {
			z = rootTracker.right
		}
		y := r.down

		for y != nil && y.down != nil && z != nil && z.down != nil {
			if y.val > z.val {
				x.down, z.down, y.down = z.down, y.down, x.down
				y = y.down
			} else if y.val < z.val {
				y = y.down
			} else {
				z = z.down
			}
		}

		rootTracker = rootTracker.right
	}

	return r.down
}

type listNodex struct {
	val   int
	right *listNodex
	down  *listNodex
}

func printList(head *listNodex) {
	for head != nil {
		fmt.Printf("%d -> ", head.val)
		head = head.down
	}
	fmt.Println()
}
func listNode_new(val int) *listNodex {
	var node *listNodex = new(listNodex)
	node.val = val
	node.right = nil
	node.down = nil
	return node
}
