package main

import "fmt"

func findParent(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = findParent(parent, parent[x])
	}
	return parent[x]
}

func unionParent(parent []int, a, b int) {
	a = findParent(parent, a)
	b = findParent(parent, b)
	if a < b {
		parent[b] = a
	} else {
		parent[a] = b
	}
}

func main() {
	var v, e int
	fmt.Scanln(&v, &e)

	parent := make([]int, v+1)

	for i := 1; i < v+1; i++ {
		parent[i] = i
	}

	for i := 0; i < e; i++ {
		var a, b int
		fmt.Scanln(&a, &b)
		unionParent(parent, a, b)
	}

	for i := 1; i < v+1; i++ {
		fmt.Printf("%d ", findParent(parent, i))
	}

	fmt.Println(parent)
}
