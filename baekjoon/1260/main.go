package main

import (
	"fmt"
)

var V, E, S int
var arr [][]int
var visited []bool

func main() {
	fmt.Scanln(&V, &E, &S)
	arr = make([][]int, V+1)
	for y := 1; y <= V; y++ {
		arr[y] = make([]int, V+1)
	}

	for y := 1; y <= E; y++ {
		var a, b int
		fmt.Scanln(&a, &b)
		arr[a][b] = 1
		arr[b][a] = 1
	}

	visited = make([]bool, V+1)
	dfs(S)
	fmt.Println()

	visited = make([]bool, V+1)
	bfs(S)
	fmt.Println()
}

func dfs(n int) {
	visited[n] = true
	fmt.Printf("%d ", n)

	for x := 1; x <= V; x++ {
		if arr[n][x] == 1 && !visited[x] {
			dfs(x)
		}
	}

}

func bfs(n int) {
	queue := []int{n}
	visited[n] = true

	for len(queue) > 0 {
		y := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", y)
		for x := 1; x <= V; x++ {
			if arr[y][x] == 1 && !visited[x] {
				queue = append(queue, x)
				visited[x] = true
			}
		}
	}
}
