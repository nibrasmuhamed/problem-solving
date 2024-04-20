package main

import (
	heaps "container/heap"
	"container/list"
	"fmt"
	"math"
)

type heap struct {
	items []int
}

func buildTree(values []int, index *int) *treeNode {
	if *index >= len(values) || values[*index] == -1 {
		*index++
		return nil
	}

	root := &treeNode{value: values[*index]}
	*index++

	root.left = buildTree(values, index)
	root.right = buildTree(values, index)

	return root
}

type treeNode struct {
	value       int
	right, left *treeNode
	next        *treeNode
}
type treeNodeDist struct {
	t *treeNode
	d int
}

func preOrderIterative(root *treeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	stack := []*treeNode{root}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, current.value)

		// Note: Right child is pushed before left to ensure left is processed first (LIFO).
		if current.right != nil {
			stack = append(stack, current.right)
		}
		if current.left != nil {
			stack = append(stack, current.left)
		}
	}
	return result
}

func inOrderIterative(root *treeNode) {
	if root == nil {
		return
	}

	stack := list.New()
	current := root

	for current != nil || stack.Len() > 0 {
		// Traverse to the leftmost node
		for current != nil {
			stack.PushBack(current)
			current = current.left
		}

		// Visit the top of the stack (current node)
		top := stack.Back()
		node := top.Value.(*treeNode)
		stack.Remove(top)
		fmt.Println(node.value)

		// Move to the right subtree
		current = node.right
	}

}
func postOrderIterative(root *treeNode) {
	if root == nil {
		return
	}

	stack := list.New()
	current := root
	var prev *treeNode
	for current != nil || stack.Len() > 0 {
		// Traverse to the leftmost node
		for current != nil {
			stack.PushBack(current)
			current = current.left
		}

		// Peek at the top of the stack without removing it
		top := stack.Back()
		node := top.Value.(*treeNode)

		// Check if the current node's right subtree needs to be processed
		if node.right == nil || node.right == prev {
			// Visit the current node
			fmt.Println(node.value)
			stack.Remove(top)
			prev = node
		} else {
			// Move to the right subtree
			current = node.right
		}
	}

}
func t2Sum(A *treeNode, B int) int {
	m := map[int]int{}
	if checkSum(A, B, m) {
		return 1
	}
	return 0
}
func checkSum(A *treeNode, B int, m map[int]int) bool {
	if A == nil {
		return false
	}
	v := A.value
	_, ok := m[B-v]
	if ok {
		return true
	}
	m[v]++
	return checkSum(A.left, B, m) || checkSum(A.right, B, m)
}
func count(A *treeNode, B int, C int) int {
	return helper(A, B, C)
}

func helper(A *treeNode, B, C int) int {
	if A == nil {
		return 0
	}
	c := 0
	if A.value >= B && A.value <= C {
		c++
	}
	l := helper(A.left, B, C)
	r := helper(A.right, B, C)

	return l + r + c
}
func solve(A []int) string {
	root := A[0]
	l, r := -1<<31, 1<<31-1
	for i := 1; i < len(A); i++ {
		if A[i] > root {
			l = root
		} else {
			r = root
		}
		if A[i] < l || A[i] > r {
			return "NO"
		}
		root = A[i]
	}
	return "YES"
}
func deleteNode(root *treeNode, key int) *treeNode {
	var parent, current *treeNode
	current = root

	for current != nil && current.value != key {
		parent = current
		if key < current.value {
			current = current.left
		} else {
			current = current.right
		}
	}

	if current == nil {
		return root // Key not found, no changes needed
	}

	// Node with only one child or no child
	if current.left == nil {
		if parent == nil {
			return current.right
		}
		if parent.left == current {
			parent.left = current.right
		} else {
			parent.right = current.right
		}
	} else if current.right == nil {
		if parent == nil {
			return current.left
		}
		if parent.left == current {
			parent.left = current.left
		} else {
			parent.right = current.left
		}
	} else {
		// Node with two children, get the in-order predecessor (largest in the left subtree)
		predecessor := findMax(current.left)
		current.value = predecessor.value
		current.left = deleteNode(current.left, predecessor.value)
	}

	return root
}

func findMax(node *treeNode) *treeNode {
	for node.right != nil {
		node = node.right
	}
	return node
}
func isValidBST(A *treeNode) int {
	x := validate(A, -1<<63, 1<<63-1)
	if x {
		return 1
	}
	return 0
}

func validate(a *treeNode, left, right int64) bool {
	if a == nil {
		return true
	}
	// Correct boundary check:
	if int64(a.value) < left || int64(a.value) >= right {
		return false
	}
	return validate(a.left, left, int64(a.value)) && validate(a.right, int64(a.value), right)
}
func buildTreeDeserialise(A []int) *treeNode {
	q := list.New()
	r := &treeNode{value: A[0]}
	q.PushBack(r)
	i := 1
	for q.Len() > 0 {
		node := q.Front()

		if node == nil {
			continue
		}

		leftVal, rightVal := -1, -1
		if i < len(A) {
			leftVal = A[i]
			i++
		}
		if i < len(A) {
			rightVal = A[i]
			i++
		}

		x := node.Value.(*treeNode)

		if leftVal != -1 {
			x.left = &treeNode{value: leftVal}
			q.PushBack(x.left)
		}

		if rightVal != -1 {
			x.right = &treeNode{value: rightVal}
			q.PushBack(x.right)
		}

		q.Remove(node)
	}
	return r
	// return buildDeserialise(A, 0)
}
func buildDeserialise(A []int, idx int) *treeNode {
	if idx > len(A) {
		return nil
	}
	if A[idx] == -1 {
		return nil
	}
	n := &treeNode{}
	n.value = A[idx]
	n.left = buildDeserialise(A, 2*idx+1)
	n.right = buildDeserialise(A, 2*idx+2)
	return n
}
func serilize(t *treeNode) []int {
	a := []int{}
	if t == nil {
		return a
	}
	q := &list.List{}
	q.PushBack(t)
	for q.Len() > 0 {
		node, x := q.Front(), q.Front().Value.(*treeNode)

		if x.left != nil && x.value != -1 {
			q.PushBack(x.left)
		} else if x.value != -1 {
			q.PushBack(&treeNode{-1, nil, nil, nil})
		}
		if x.right != nil && x.value != -1 {
			q.PushBack(x.right)
		} else if x.value != -1 {
			q.PushBack(&treeNode{-1, nil, nil, nil})
		}

		a = append(a, x.value)

		q.Remove(node)
	}
	// a = append(a, q.Back().Value.(*treeNode).value)
	return a
}
func zigzagLevelOrder(t *treeNode) [][]int {
	a := [][]int{}
	if t == nil {
		return [][]int{}
	}
	y := t
	q := &list.List{}
	q.PushBack(t)
	zig := false
	j := []int{}
	for q.Len() > 0 {
		node, x := q.Front(), q.Front().Value.(*treeNode)
		j = append(j, x.value)
		if zig {
			if x.left != nil {
				q.PushBack(x.left)
			}
			if x.right != nil {
				q.PushBack(x.right)
			}
		} else {
			if x.right != nil {
				q.PushBack(x.right)
			}
			if x.left != nil {
				q.PushBack(x.left)
			}
		}

		if y == x && q.Len() > 0 {
			a = append(a, j)
			j = []int{}
			y = q.Back().Value.(*treeNode)
		}
		q.Remove(node)
		zig = !zig
	}
	return a
}
func verticalOrderTraversal(t *treeNode) [][]int {
	a := [][]int{}
	if t == nil {
		return [][]int{}
	}
	n := &treeNodeDist{t, 0}
	minDist, maxDist := math.MaxInt, math.MinInt
	m := map[int][]int{}
	q := &list.List{}
	q.PushBack(n)
	// j := []int{}
	for q.Len() > 0 {
		node, x := q.Front(), q.Front().Value.(*treeNodeDist)
		// j = append(j, x.value)
		_, ok := m[x.d]
		if !ok {
			m[x.d] = []int{x.t.value}
		} else {
			m[x.d] = append(m[x.d], x.t.value)
		}
		q.Remove(node)
		if x.t.left != nil {
			leftQueue := &treeNodeDist{x.t.left, x.d - 1}
			if x.d < minDist {
				minDist = x.d
			}
			q.PushBack(leftQueue)
		}
		if x.t.right != nil {
			rightQueue := &treeNodeDist{x.t.right, x.d + 1}
			if x.d > maxDist {
				maxDist = x.d
			}
			q.PushBack(rightQueue)
		}
	}
	// a = append(a, j)
	for i := minDist - 1; i <= maxDist+1; i++ {
		a = append(a, m[i])
	}
	return a
}
func topView(t *treeNode) []int {
	a := []int{}
	if t == nil {
		return []int{}
	}
	n := &treeNodeDist{t, 0}
	minDist, maxDist := 0, 0
	m := map[int][]int{}
	q := &list.List{}
	q.PushBack(n)
	// j := []int{}
	for q.Len() > 0 {
		node, x := q.Front(), q.Front().Value.(*treeNodeDist)
		// j = append(j, x.value)
		_, ok := m[x.d]
		if !ok {
			m[x.d] = []int{x.t.value}
		} else {
			m[x.d] = append(m[x.d], x.t.value)
		}

		if x.t.left != nil {
			leftQueue := &treeNodeDist{x.t.left, x.d - 1}
			if x.d-1 < minDist {
				minDist = x.d - 1
			}
			q.PushBack(leftQueue)
		}
		if x.t.right != nil {
			rightQueue := &treeNodeDist{x.t.right, x.d + 1}
			if x.d+1 > maxDist {
				maxDist = x.d + 1
			}
			q.PushBack(rightQueue)
		}
		q.Remove(node)
	}
	// a = append(a, j)
	for i := minDist; i <= maxDist; i++ {
		a = append(a, m[i][0])
	}
	return a
}
func rightView(t *treeNode) []int {
	a := []int{}
	if t == nil {
		return []int{}
	}
	y := t
	q := &list.List{}
	q.PushBack(t)
	for q.Len() > 0 {
		node, x := q.Front(), q.Front().Value.(*treeNode)
		if x.left != nil {
			q.PushBack(x.left)
		}
		if x.right != nil {
			q.PushBack(x.right)
		}
		if y == x && q.Len() > 0 {
			a = append(a, x.value)
			y = q.Back().Value.(*treeNode)
		}
		q.Remove(node)
	}
	return a
}
func levelOrder(t *treeNode) [][]int {
	a := [][]int{}
	if t == nil {
		return [][]int{}
	}
	y := t
	q := &list.List{}
	q.PushBack(t)
	j := []int{}
	for q.Len() > 0 {
		node, x := q.Front(), q.Front().Value.(*treeNode)
		j = append(j, x.value)
		q.Remove(node)
		if x.left != nil {
			q.PushBack(x.left)
		}
		if x.right != nil {
			q.PushBack(x.right)
		}
		if y == x && q.Len() > 0 {
			// fmt.Println()
			a = append(a, j)
			j = []int{}
			y = q.Back().Value.(*treeNode)
		}
	}
	a = append(a, j)
	return a
}
func listEg() {
	q := &list.List{}
	q.PushBack(8)
	s := q.Front().Value.(int)
	fmt.Println(s)
	// b := q.Front()
	// q.Remove(b)
	fmt.Printf("q: %v\n", q.Front().Next())
}
func morrisInorder(A *treeNode) []int {
	root := A
	out := []int{}
	for root != nil {
		if root.left == nil {
			out = append(out, root.value)
			root = root.right
		} else {
			prev := root.left
			for prev.right != nil && prev.right != root {
				prev = prev.right
			}

			if prev.right == nil {
				prev.right = root
				root = root.left
			} else {
				prev.right = nil
				out = append(out, root.value)
				root = root.right
			}
		}
	}
	return out
}

func kthsmallest(A *treeNode, B int) int {
	root := A
	counter := 0
	for root != nil {
		if root.left == nil {
			// out = append(out, root.value)
			counter++
			if counter == B {
				return root.value
			}
			root = root.right
		} else {
			prev := root.left
			for prev.right != nil && prev.right != root {
				prev = prev.right
			}

			if prev.right == nil {
				prev.right = root
				root = root.left
			} else {
				prev.right = nil
				counter++
				if counter == B {
					return root.value
				}
				root = root.right
			}
		}
	}
	return -1
}

func findTwoVals(a []int) []int {
	low, high := -1, -1

	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			if low == -1 {
				low = a[i]
				high = a[i+1]
			} else {
				high = a[i+1]
			}
		}
	}

	return []int{low, high}
}

func inorderTraversal(node *treeNode, count *int, result *int, k int) {
	if node == nil || *count >= k {
		return
	}

	inorderTraversal(node.left, count, result, k)
	*count++

	if *count == k {
		*result = node.value
		return
	}

	inorderTraversal(node.right, count, result, k)
}

func kthSmallest(root *treeNode, k int) int {
	count := 0
	result := 0

	// Start the recursive inorder traversal from the root
	inorderTraversal(root, &count, &result, k)

	return result
}

func findPath(root *treeNode, k int, path *[]int) bool {
	if root == nil {
		return false
	}
	*path = append(*path, root.value)

	if k == root.value {
		return true
	}
	if findPath(root.left, k, path) || findPath(root.right, k, path) {
		return true
	}
	*path = (*path)[:len(*path)-1]
	return false
}

func equalSum(A *treeNode) int {
	totalSum := sumTree(A)
	if equalSumCheck(A, totalSum/2) {
		return 1
	}
	return 0
}
func equalSumCheck(A *treeNode, targetSum int) bool {
	if A == nil {
		return false
	}
	if equalSumCheck(A.left, targetSum) || equalSumCheck(A.right, targetSum) {
		return true
	}
	leftSum := sumTree(A.left)
	rightSum := sumTree(A.right)
	return leftSum+rightSum+A.value == targetSum
}
func sumTree(A *treeNode) int {
	if A == nil {
		return 0
	}
	return sumTree(A.left) + sumTree(A.right) + A.value
}

func connect(root *treeNode) *treeNode {
	A := root
	curr, last, temp := root, root, root
	for curr != nil {
		if curr.left != nil {
			temp.next = curr.left
			temp = temp.next
		}
		if curr.right != nil {
			temp.next = curr.right
			temp = temp.next
		}
		if curr == last {
			curr = last.next
			last.next = nil
			last = temp
		} else {
			curr = curr.next
		}
	}
	return A
}
func hasPathSum(A *treeNode, B int) int {
	if A == nil {
		return 0
	}
	if A.left == nil && A.right == nil && A.value == B {
		return 1
	}
	x := hasPathSum(A.left, B-A.value)
	y := hasPathSum(A.right, B-A.value)
	if x == 1 || y == 1 {
		return 1
	} else {
		return 0
	}
}
func checkPathVal(root *treeNode, b int) bool {
	if root == nil {
		return false
	}
	if root.left == nil && root.right == nil {
		return root.value == b
	}
	return checkPathVal(root.left, b-root.value) || checkPathVal(root.right, b-root.value)
}
func distance(root *treeNode, b, c int) int {
	var left, right int
	// var temp *treeNode
	for root != nil {
		if b > root.value && c < root.value {
			right = calcDistance(root, b)
			left = calcDistance(root, c)
			return left + right
		} else if b < root.value && c > root.value {
			right = calcDistance(root, c)
			left = calcDistance(root, b)
			return left + right
		} else if b == root.value {
			if c > root.value {
				right = calcDistance(root, c)
			} else {
				left = calcDistance(root, c)
			}
			return left + right
		} else if c == root.value {
			if b > root.value {
				right = calcDistance(root, b)
			} else {
				left = calcDistance(root, b)
			}
			return left + right
		}
		if b > root.value && c > root.value {
			root = root.right
		} else if b < root.value && c < root.value {
			root = root.left
		}
	}
	return 0
}

func calcDistance(root *treeNode, b int) int {
	d := 0
	for root != nil {
		if root.value == b {
			return d
		}
		d++
		if root.value > b {
			root = root.left
		} else {
			root = root.right
		}
	}
	return d
}

func invertTree(t *treeNode) *treeNode {
	if t == nil {
		return nil
	}
	t.left, t.right = t.right, t.left
	t.left = invertTree(t.left)
	t.right = invertTree(t.right)
	return t
}

func identicalTree(a, b *treeNode) bool {
	if a == nil && b == nil {
		return true
	}

	// If one of the nodes is nil and the other is not,
	// they are not equal
	if a == nil || b == nil {
		return false
	}
	l := identicalTree(a.left, b.left)
	r := identicalTree(a.right, b.right)
	return l && r && (a.value == b.value)
}

func partition(t *treeNode, m *map[int]bool, x *bool) int {
	if t == nil {
		return 0
	}
	leftSum := partition(t.left, m, x)
	rightSum := partition(t.right, m, x)

	_, ok := (*m)[leftSum]
	if ok {
		*x = true
	}

	_, ok = (*m)[rightSum]
	if ok {
		*x = true
	}

	(*m)[leftSum] = true
	(*m)[rightSum] = true

	return leftSum + rightSum + t.value
}

func heapify(A []int) []int {
	n := len(A)

	for i := n/2 - 1; i >= 0; i-- {
		minHeapify(A, i, n)
	}

	return A
}

func minHeapify(A []int, i, n int) {
	for {
		left, right := 2*i+1, 2*i+2
		smallest := i

		if left < n && A[left] < A[smallest] {
			smallest = left
		}

		if right < n && A[right] < A[smallest] {
			smallest = right
		}

		if smallest != i {
			A[i], A[smallest] = A[smallest], A[i]
			i = smallest
		} else {
			break
		}
	}
}
func connectRopes(A []int) int {
	h := &heap{}
	for _, v := range A {
		h.insert(v)
	}
	cost := 0
	for !h.isEmpty() {
		r1 := h.get()    // Get only one element
		if h.isEmpty() { // Check for only one element remaining
			break // Exit the loop if only one element is left
		}
		r2 := h.get() // Get the second element
		x := r1 + r2
		cost += x
		h.insert(x)
	}
	return cost
}
func (h *heap) isEmpty() bool {
	return len(h.items) == 0
}
func (h *heap) get() int {
	if len(h.items) == 0 {
		return -1
	}

	val := h.items[0]
	lastIdx := len(h.items) - 1

	h.items[0] = h.items[lastIdx]
	h.items = h.items[:lastIdx]

	i := 0
	for i*2+1 < len(h.items) {
		left, right := i*2+1, i*2+2
		smallest := left

		if right < len(h.items) && h.items[right] < h.items[left] {
			smallest = right
		}

		if h.items[i] > h.items[smallest] {
			h.items[i], h.items[smallest] = h.items[smallest], h.items[i]
			i = smallest
		} else {
			break
		}
	}

	return val
}
func (h *heap) insert(val int) {
	h.items = append(h.items, val)
	for i := len(h.items) - 1; i >= 0; {
		p := (i - 1) / 2
		if h.items[p] > h.items[i] {
			h.items[(i-1)/2], h.items[i] = h.items[i], h.items[(i-1)/2]
			i = p
		} else {
			break
		}
	}
}

func checkHeap(items []int) bool {
	i := 0
	for i = 0; i*2+2 < len(items); i++ {
		if items[i] > items[i*2+1] || items[i] > items[i*2+2] {
			return false
		}
	}
	if i*2+1 < len(items) && items[i] > items[i*2+1] {
		return false
	}
	return true
}

type ListNode struct {
	value int
	next  *ListNode
}

type MinHeap []*ListNode

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func mergeKLists(lists []*ListNode) *ListNode {
	minHeap := &MinHeap{}
	for i := 0; i < len(lists); i++ {
		heaps.Push(minHeap, lists[i])
	}

	dummyHead := &ListNode{0, nil}
	current := dummyHead

	for minHeap.Len() > 0 {
		node := heaps.Pop(minHeap).(*ListNode)
		current.next = node
		current = current.next

		if node.next != nil {
			heaps.Push(minHeap, node.next)
		}
	}

	return dummyHead.next
}

type MaxHeapInt []int

func (h MaxHeapInt) Len() int           { return len(h) }
func (h MaxHeapInt) Less(i, j int) bool { return h[i] > h[j] }
func (h *MaxHeapInt) Swap(i, j int) {
	if len(*h) > 0 {
		(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
	}
}

func (h *MaxHeapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeapInt) Pop() interface{} {
	if len(*h) == 0 {
		return -1
	}
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MinHeapInt []int

func (h MinHeapInt) Len() int           { return len(h) }
func (h MinHeapInt) Less(i, j int) bool { return h[i] < h[j] }
func (h *MinHeapInt) Swap(i, j int) {
	if len(*h) > 0 {
		(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
	}
}

func (h *MinHeapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeapInt) Pop() interface{} {
	if len(*h) == 0 {
		return -1
	}
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func largestProductOfThreeValues(A []int) []int {
	out := make([]int, 0)
	maxheap := &MaxHeapInt{}
	for i := 0; i < 2; i++ {
		heaps.Push(maxheap, A[i])
		out = append(out, -1)
	}
	for i := 2; i < len(A); i++ {
		heaps.Push(maxheap, A[i])
		val1, val2, val3 := heaps.Pop(maxheap).(int), heaps.Pop(maxheap).(int), heaps.Pop(maxheap).(int)
		out = append(out, val1*val2*val3)
		heaps.Push(maxheap, val1)
		heaps.Push(maxheap, val2)
		heaps.Push(maxheap, val3)
	}
	return out
}
func runningMedian(A []int) []int {
	mxHeap, mnHeap := &MaxHeapInt{}, &MinHeapInt{}
	res := make([]int, 0)
	heaps.Init(mnHeap)
	heaps.Init(mnHeap)
	for i := 0; i < len(A); i++ {
		if mxHeap.Len() > 0 && (*mxHeap)[0] <= A[i] {
			heaps.Push(mnHeap, A[i])
		} else {
			heaps.Push(mxHeap, A[i])
		}
		diff := mxHeap.Len() - mnHeap.Len()
		if diff > 1 {
			value := heaps.Pop(mxHeap).(int)
			heaps.Push(mnHeap, value)
		} else if diff < 0 {
			value := heaps.Pop(mnHeap).(int)
			heaps.Push(mxHeap, value)
		}
		res = append(res, (*mxHeap)[0])
	}
	return res
}

func main() {
	// list1 := &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
	// list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	// list3 := &ListNode{2, &ListNode{6, nil}}

	// // Merging the linked lists
	// mergedList := mergeKLists([]*ListNode{list1, list2, list3})
	// fmt.Println(mergedList)
	// root := &treeNode{value: 5}
	// root.left = &treeNode{value: 3}
	// root.left.left = &treeNode{value: 1}
	// root.left.right = &treeNode{value: 4}

	// root.right = &treeNode{value: 6}
	// root.right.right = &treeNode{value: 8}

	// root := &treeNode{value: 6}
	// root.left = &treeNode{value: 9}
	// root.right = &treeNode{value: 4}
	// root.left.left = nil
	// root.left.right = nil
	// root.right.left = &treeNode{value: 8}
	// root.right.left.right = &treeNode{value: 3}
	// a := []int{1, 2, 4, -1, 3, -1, 5, 7, -1, -1, 6, -1, 8, -1, -1, -1, -1}
	// var x bool
	// m := make(map[int]bool)
	// partition(root, &m, &x)
	// fmt.Println(equalSum(root))
	// a := []int{63, 23, 64, 1, 67, 12, 24, 86}
	// fmt.Println(sortWithHeap(a))
	// // fmt.Println(h.get())
	// fmt.Println(checkHeap(a))
	// // fmt.Println(connectRopes([]int{5, 12, 16, 5, 12, 19, 4, 3, 18, 15, 16}))
	// heapify(a)
	// A := []int{3, 2, 6}
	// B := []int{9, 8, 9}
	// (sortByA(A, B))
	// fmt.Println(A)
	// fmt.Println(B)
	// fmt.Println(findMinimumNoOfCoins(47))
	// fmt.Println(findTwoVals([]int{1, 5, 3, 4, 2, 6}))
}

func sortPartiallySorted(A []int, B int) []int {
	h := &MinHeapInt{}
	heaps.Init(h)
	for i := 0; i <= B; i++ {
		heaps.Push(h, A[i])
	}
	out := []int{}
	for i := B + 1; i < len(A) || h.Len() > 0; i++ {
		out = append(out, heaps.Pop(h).(int))
		if i < len(A) {
			heaps.Push(h, A[i])
		}
	}
	return out
}

func kthLargest(A int, B []int) []int {
	out := []int{}
	for i := 0; i < A-1; i++ {
		out = append(out, -1)
	}
	h := &MinHeapInt{}
	heaps.Init(h)
	for i := 0; i < A; i++ {
		heaps.Push(h, B[i])
	}
	out = append(out, (*h)[0])
	for i := A; i < len(B); i++ {
		heaps.Push(h, B[i])
		heaps.Pop(h)
		out = append(out, (*h)[0])
	}
	return out
}
func extractFromHeap(A [][]int) []int {
	h := &MinHeapInt{}
	heaps.Init(h)
	out := []int{}
	for _, v := range A {
		if v[0] == 1 {
			out = append(out, heaps.Pop(h).(int))
		} else {
			heaps.Push(h, v[1])
		}
	}
	return out
}

// type maxHeap []int

func sortWithHeap(A []int) []int {
	n := len(A)
	heapifyTill(A, 0, n)
	for i := n - 1; i > 0; i-- {
		// Swap the root (maximum element) with the last element
		A[0], A[i] = A[i], A[0]

		// Call maxHeapify on the reduced heap
		maxHeapify(A, 0, i)
	}
	// A[j], A[0] = A[0], A[j]
	return A
}
func heapifyTill(A []int, start, end int) {
	for i := (end / 2) - 1; i >= 0; i-- {
		maxHeapify(A, i, end)
	}
}
func maxHeapify(A []int, start, end int) {
	for {
		left, right := start*2+1, start*2+2
		largest := start

		if left < end && A[left] > A[largest] {
			largest = left
		}
		if right < end && A[right] > A[largest] {
			largest = right
		}
		if start != largest {
			A[start], A[largest] = A[largest], A[start]
			start = largest
		} else {
			break
		}
	}
}

func heapifyUp(A []int, index int) {
	for index > 0 {
		parent := (index - 1) / 2

		if A[index] > A[parent] {
			// Swap with parent if the current node is greater
			A[index], A[parent] = A[parent], A[index]
			index = parent
		} else {
			break // Break if the heap property is restored
		}
	}
}

// Heap StartsFrom Here!

func Heapify(arr []int) {
	n := len(arr)

	// Heapify down example
	for i := n/2 - 1; i >= 0; i-- {
		HeapifyDown(arr, i, n)
	}
}
func HeapifyDown(A []int, start, end int) {
	for {
		left, right := start*2+1, start*2+2
		largest := start

		if left < end && A[left] > A[largest] {
			largest = left
		}
		if right < end && A[right] > A[largest] {
			largest = right
		}
		if start != largest {
			A[start], A[largest] = A[largest], A[start]
			start = largest
		} else {
			break
		}
	}
}

// HeapifyUp performs heapify up operation on a max heap
func HeapifyUp(A []int, start int) {
	for start > 0 {
		parent := (start - 1) / 2
		if A[start] > A[parent] {
			A[start], A[parent] = A[parent], A[start]
			start = parent
		} else {
			break
		}
	}
}
