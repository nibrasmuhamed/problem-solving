package main

import "fmt"

type Graph map[int][]int

func (g *Graph) print() {
	for i, val := range *g {
		fmt.Println(i, " is ", val)
	}
}

func (g *Graph) addVertex(data int) {
	x := []int{}
	(*g)[data] = x
}

func (g *Graph) insert(vertex, edge int, isBidirectional bool) {
	_, ok := (*g)[vertex]
	if !ok {
		g.addVertex(vertex)
	}
	_, ok = (*g)[edge]
	if !ok {
		g.addVertex(edge)
	}
	(*g)[vertex] = append((*g)[vertex], edge)
	if isBidirectional {
		(*g)[edge] = append((*g)[edge], vertex)
	}
}

func main() {
	g := Graph{}
	g.insert(3, 4, true)
	g.print()
}
