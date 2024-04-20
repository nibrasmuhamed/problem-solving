package main

import (
	"fmt"
)

type Stack struct {
	items []int
}

// Push adds an item to the top of the stack
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop removes and returns the item from the top of the stack
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		// Handle empty stack case, returning a default value (you can choose a different approach)
		return -1
	}
	popped := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return popped
}

// Peek returns the item from the top of the stack without removing it
func (s *Stack) Peek() int {
	if len(s.items) == 0 {
		// Handle empty stack case, returning a default value (you can choose a different approach)
		return -1
	}
	return s.items[len(s.items)-1]
}

// IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
func sortStackWithStack(A []int) []int {
	s1 := Stack{items: A}
	s2 := Stack{items: []int{}}

	for !s1.IsEmpty() {
		x := s1.Pop()
		for !s2.IsEmpty() && x < s2.Peek() {
			temp := s2.Pop()
			s1.Push(temp)
		}
		s2.Push(x)
	}
	return s2.items
}

func main() {
	fmt.Println(leftGreater([]int{2, 1, 3, 2, 1, 2, 4, 3, 2, 1, 3, 1}))
}

// neareast smaller element on left
func leftSmaller(A []int) []int {
	out := []int{}
	stack := &Stack{}
	for i := range A {
		for !stack.IsEmpty() && A[stack.Peek()] > A[i] {
			stack.Pop()
		}
		if stack.IsEmpty() {
			out = append(out, -1)
		} else {
			out = append(out, stack.Peek())
		}
		stack.Push(i)
	}
	return out
}
func leftGreater(A []int) []int {
	out := []int{}
	stack := &Stack{}
	for i := range A {
		for !stack.IsEmpty() && A[stack.Peek()] < A[i] {
			stack.Pop()
		}
		if stack.IsEmpty() {
			out = append(out, -1)
		} else {
			out = append(out, stack.Peek())
		}
		stack.Push(i)
	}
	return out
}
func rightSmaller(A []int) []int {
	out := make([]int, len(A))
	stack := &Stack{}
	for i := len(A) - 1; i >= 0; i-- {
		for !stack.IsEmpty() && A[stack.Peek()] >= A[i] {
			stack.Pop()
		}
		if stack.IsEmpty() {
			out[i] = -1
		} else {
			out[i] = stack.Peek()
		}
		stack.Push(i)
	}
	return out
}
func rightGreater(A []int) []int {
	out := make([]int, len(A))
	stack := &Stack{}
	for i := len(A) - 1; i >= 0; i-- {
		for !stack.IsEmpty() && A[stack.Peek()] <= A[i] {
			stack.Pop()
		}
		if stack.IsEmpty() {
			out[i] = -1
		} else {
			out[i] = stack.Peek()
		}
		stack.Push(i)
	}
	return out
}
func area(A []int) int {
	rightGreat, rightSmall := rightGreater(A), rightSmaller(A)
	leftGreat, leftSmall := leftGreater(A), leftSmaller(A)
	sum := 0
	for i := range A {
		subArrayCountMax := (i - leftGreat[i]) * (rightGreat[i] - i)
		subArrayCountMin := (i - leftSmall[i]) * (rightSmall[i] - i)
		sum += A[i] * (subArrayCountMax - subArrayCountMin)
	}
	return sum
}

func largestRectangleArea(A []int) int {
	left, right := leftSmaller(A), rightSmaller(A)
	if len(A) == 1 {
		return A[0] * 1
	}
	ans := 0
	for i, v := range A {
		j := left[i]
		k := right[i]
		// if j == -1 {
		// 	j = 0
		// }
		if k == -1 {
			k = len(A)
		}
		ans = max(ans, v*(k-j-1))
		// fmt.Println("a[", i, "] = ", A[i], " k = ", k, " j = ", j, " total = ", v*(k-j-1))
	}
	return ans
}

// func braces(A string) int {
// 	stack := Stack{}
// 	for _, v := range A {
// 		if string(v) == "+" || string(v) == "-" || string(v) == "/" || string(v) == "*" || string(v) == "(" {
// 			stack.Push(string(v))
// 		} else if string(v) == ")" {
// 			operator := 0
// 			for !stack.IsEmpty() {
// 				if stack.Peek() != "(" {
// 					operator++
// 				}
// 				if stack.Peek() == "(" && operator == 0 {
// 					return 1
// 				}
// 				q := stack.Pop()
// 				if q == "(" {
// 					break
// 				}
// 			}
// 		}
// 	}
// 	return 0
// }

// func evalRPN(A []string) int {
// 	stack := Stack{}
// 	sum := 0
// 	for _, v := range A {
// 		if v == "+" || v == "-" || v == "/" || v == "*" {
// 			v1 := stack.Pop()
// 			v2 := stack.Pop()
// 			x := evalute(v, v2, v1)
// 			stack.Push(x)
// 		} else {
// 			stack.Push(v)
// 		}
// 	}
// 	sum, _ = strconv.Atoi(stack.Pop())
// 	return sum
// }
// func evalute(v, v2, v1 string) string {
// 	x, _ := strconv.Atoi(v2)
// 	y, _ := strconv.Atoi(v1)
// 	switch v {
// 	case "+":
// 		return fmt.Sprint(x + y)
// 	case "-":
// 		return fmt.Sprint(x - y)
// 	case "/":
// 		return fmt.Sprint(x / y)
// 	case "*":
// 		return fmt.Sprint(x * y)
// 	}
// 	return ""
// }

// func stringbalance(A string) string {
// 	s := Stack{}

// 	for i := 0; i < len(A); i++ {
// 		peekedVal := s.Peek()
// 		if peekedVal == string(A[i]) {
// 			s.Pop()
// 		} else {
// 			s.Push(string(A[i]))
// 		}
// 	}
// 	str := ""
// 	for !s.IsEmpty() {
// 		str += s.Pop() // Add space after each character to separate words
// 	}
// 	return reverse(str)
// }
// func reverse(s string) string {
// 	runes := []rune(s)

// 	// Get length of the rune slice
// 	n := len(runes)

// 	// Swap elements in the slice
// 	for i := 0; i < n/2; i++ {
// 		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
// 	}

// 	// Convert back to string
// 	return string(runes)
// }

// func Passes(A int, B int, C []int) int {
// 	stack := Stack{}
// 	val := B
// 	for i := 0; i < A; i++ {
// 		if C[i] != 0 {
// 			if i == 0 {
// 				stack.Push(B)
// 			} else {
// 				stack.Push(val)
// 			}
// 			val = C[i]
// 		} else {
// 			val = stack.Pop()
// 		}
// 	}
// 	return val
// }

// func BalancedParenthesis(A string) int {
// 	s := Stack{}
// 	for i := 0; i < len(A); i++ {
// 		if string(A[i]) == "(" || string(A[i]) == "[" || string(A[i]) == "{" {
// 			s.Push(string(A[i]))
// 		} else {
// 			v := s.Pop()
// 			if string(A[i]) == ")" && v != "(" {
// 				return 0
// 			}
// 			if string(A[i]) == "]" && v != "[" {
// 				return 0
// 			}
// 			if string(A[i]) == "}" && v != "{" {
// 				return 0
// 			}
// 		}
// 	}
// 	if !s.IsEmpty() {
// 		return 0
// 	}
// 	return 1
// }
