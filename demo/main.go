package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	var y int = 1<<31 - 1
	y += 2
	fmt.Println(reflect.TypeOf(y))
	fmt.Println(y)
	fmt.Println(math.MaxInt32)
	// fmt.Println(math.MinInt64)
}
