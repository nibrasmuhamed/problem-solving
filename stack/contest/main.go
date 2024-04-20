package main

import (
	"fmt"
	"math"
	"sort"
)

func distanceToOrigin(x, y, px, py int) float64 {
	return math.Sqrt(float64((x-px)*(x-px) + (y-py)*(y-py)))
}

func closestPointToOrigin(points [][]int) []int {
	closest := points[0]
	minDistance := distanceToOrigin(closest[0], closest[1], 0, 0)

	for i := 1; i < len(points); i++ {
		currentPoint := points[i]
		currentDistance := distanceToOrigin(currentPoint[0], currentPoint[1], 0, 0)

		// Check if the current point is closer or has the same distance with smaller x-coordinate
		if currentDistance < minDistance || (currentDistance == minDistance && currentPoint[0] < closest[0]) {
			closest = currentPoint
			minDistance = currentDistance
		}
	}

	return closest
}

func sortByDistance(points [][]int, p []int) {
	sort.Slice(points, func(i, j int) bool {
		distanceI := distanceToOrigin(points[i][0], points[i][1], p[0], p[1])
		distanceJ := distanceToOrigin(points[j][0], points[j][1], p[0], p[1])

		// Sort by distance, then by x, and then by y
		if distanceI != distanceJ {
			return distanceI < distanceJ
		}
		if points[i][0] != points[j][0] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] < points[j][1]
	})
}

func findMissingNumbers(A []int, B int) int {
	m := map[int]int{}
	for i, v := range A {
		m[v] = i
	}
	out := []int{}
	for i := 1; i <= A[len(A)-1]+B; i++ {
		_, ok := m[i]
		if !ok {
			out = append(out, i)
		}
	}
	return out[B-1]
}
func countMissingNumbers(arr []int, mid int) int {
	count := 0
	for _, num := range arr {
		if num <= mid {
			count++
		} else {
			break
		}
	}
	return mid - count
}

func findMissingNumber(arr []int, B int) int {
	low, high := 1, int(1e9)
	result := -1

	for low <= high {
		mid := low + (high-low)/2

		missingCount := countMissingNumbers(arr, mid)

		if missingCount < B {
			low = mid + 1
		} else {
			result = mid
			high = mid - 1
		}
	}

	return result
}
func hasZeroSumSubarray(A []int) int {
	// Create a set to store the running sum
	sumSet := make(map[int]bool)

	// Initialize the running sum and iterate through the array
	sum := 0
	for _, num := range A {
		sum += num

		// If the current sum is zero or has been seen before, return 1
		if sum == 0 || sumSet[sum] {
			return 1
		}

		// Add the current sum to the set
		sumSet[sum] = true
	}

	// No subarray with sum zero found
	return 0
}
func solve(A []int, B int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(A))) // Sort weapons in descending order

	n := len(A)

	var moves int
	prevWeapon := -1 // Index of the previously used weapon

	for B > 0 {
		// Find the highest damage weapon that wasn't used last
		for i := 0; i < n; i++ {
			if i != prevWeapon {
				fmt.Println(A[i])
				B -= A[i]
				moves++
				prevWeapon = i
				break // Move to the next weapon in the next iteration
			}
		}
	}

	return moves
}
func main() {
	a := []int{96, 35, 48, 92, 63, 25, 61, 49, 58, 36, 80, 98, 96, 71, 19, 29, 85, 83, 40, 17, 58, 49, 12, 76, 50, 97, 17, 47, 82, 25, 53, 79,
		93, 27, 82, 93, 95, 62, 50, 40, 83, 26, 62, 24, 26, 84, 74}
	// a := []int{97, 77, 15, 76, 50, 97, 96, 20, 14, 70, 31}
	fmt.Println(solve(a, 293))
	// fmt.Println(findMissingNumber([]int{1, 2, 3, 4}, 2))
	// points := [][]int{
	// 	{1, 2},
	// 	{3, 4},
	// 	{5, 6},
	// 	{-1, -2},
	// }

	// closestPoint := closestPointToOrigin(points)
	// fmt.Printf("Closest point to the origin: (%d, %d)\n", closestPoint[0], closestPoint[1])

	// sortByDistance(points, closestPoint)
	// fmt.Println("Sorted array based on distance from closest point to the origin:", points)
}
