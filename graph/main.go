package main

import "fmt"

func main() {
	A := 5
	B := [][]int{
		{1, 2},
		{4, 1},
		{2, 4},
		{3, 4},
		{5, 2},
		{1, 3},
	}
	fmt.Println(cycleDetect(A, B))
}
func surroundedXandO(board [][]byte) {
	rows := len(board)
	cols := len(board[0])

	// capture unsurrounded region.
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// Check if the current element is on the border
			if i == 0 || i == rows-1 || j == 0 || j == cols-1 {
				surroundedXandOCaptureHelper(board, i, j)
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'T' {
				board[i][j] = 'O'
			}
		}
	}
}
func surroundedXandOCaptureHelper(board [][]byte, a, b int) {
	if a < 0 || b < 0 || a >= len(board) || b >= len(board[0]) || board[a][b] != 'O' {
		return
	}
	board[a][b] = 'T'
	surroundedXandOCaptureHelper(board, a+1, b)
	surroundedXandOCaptureHelper(board, a-1, b)
	surroundedXandOCaptureHelper(board, a, b+1)
	surroundedXandOCaptureHelper(board, a, b-1)
}

// func findShotestPathForAllNode(A int, B [][]int, C int) []int {
// 	visited := make([]bool, A)

// 	for i := 0; i < A; i++ {

// 	}
// }

type pair struct {
	from   int
	to     int
	weight int
}

func shotestPath(A int, B [][]int) int {
	// Build adjacency list from the given edges
	graph := make(map[int][]int)
	for _, edge := range B {
		from, to := edge[0], edge[1]
		graph[from] = append(graph[from], to)
	}

	// Perform DFS to check if there is a path from 1 to A
	visited := make(map[int]bool)
	return dfsShortestPath(1, A, graph, visited)
}

func dfsShortestPath(current, destination int, graph map[int][]int, visited map[int]bool) int {
	if current == destination {
		return 1
	}

	visited[current] = true

	for _, neighbor := range graph[current] {
		if !visited[neighbor] && dfsShortestPath(neighbor, destination, graph, visited) == 1 {
			return 1
		}
	}

	return 0
}
func cycleDetect(A int, B [][]int) int {
	visited := make([]bool, A+1)
	path := make([]bool, A+1)
	res := false
	for i := 1; i <= A; i++ {
		if !visited[i] {
			res = dfsCycle(i, B, path, visited)
		}
		if res {
			return 1
		}
	}
	return 0
}

func dfsCycle(u int, B [][]int, path, visited []bool) bool {
	visited[u] = true
	path[u] = true

	for _, edge := range B {
		if edge[0] == u {
			if path[edge[1]] {
				return true
			}
			if !visited[edge[1]] && dfsCycle(edge[1], B, path, visited) {
				return true
			}
		}
	}

	path[u] = false
	return false
}

func printAll(A int, B [][]int) {
	visited := make([]bool, A+1)
	for i := 1; i <= A; i++ {
		if !visited[i] {
			printAllHelper(i, B, visited)
		}
	}
}

func printAllHelper(node int, B [][]int, visited []bool) {
	visited[node] = true
	fmt.Println(node)
	for _, edge := range B {
		if edge[0] == node && !visited[edge[1]] {
			printAllHelper(edge[1], B, visited)
		}
	}
}
