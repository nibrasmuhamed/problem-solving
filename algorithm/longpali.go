package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
}

func longestPalindrome(s string) int {
	length := 0
	z := false
	hash := make(map[string]int)
	for _, val := range s {
		_, ok := hash[string(val)]
		if !ok {
			hash[string(val)] = 1
		} else {
			hash[string(val)] = hash[string(val)] + 1
		}
	}
	for _, val := range hash {
		if val%2 == 0 {
			length = length + val
			continue
		}
		if val%2 != 0 && !z {
			length = length + 1
			z = true
		}
		if val%2 != 1 && val%2 == 0 {
			length = length + (val - 1)
		}
	}
	return length
}
