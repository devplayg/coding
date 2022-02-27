package main

import (
	"fmt"
)

var vertex, edge, start int

const MAX = 1000 + 1

var arr = [MAX][MAX]int{}
var bfsVisited = [MAX]int{}
var dfsVisited = [MAX]int{}

func main() {
	fmt.Scanln(&vertex, &edge, &start)
	//println(vertex)
	//println(edge)
	//println(start)
	for i := 1; i <= edge; i++ {
		//println(i)
		var from, to int
		fmt.Scanln(&from, &to)
		arr[from][to] = 1
		arr[to][from] = 1
	}
	dfs(start)
	fmt.Println("")
	bfs(start)
	fmt.Println("")

}

func bfs(start int) {
	queue := make([]int, 0)
	queue = append(queue, start)
	bfsVisited[start] = 1

	for len(queue) > 0 {
		y := queue[0]
		fmt.Printf("%d ", y)
		queue = queue[1:]

		for x := 1; x <= vertex; x++ {
			if arr[y][x] == 1 && bfsVisited[x] == 0 {
				queue = append(queue, x)
				bfsVisited[x] = 1
			}
		}
	}
}

var dfsQueue = make([]int, 0)

func dfs(start int) {
	dfsQueue = append(dfsQueue, start)
	for len(dfsQueue) > 0 {
		_dfs()
	}

}

func _dfs() {
	y := dfsQueue[0]
	dfsVisited[y] = 1
	fmt.Printf("%d ", y)
	dfsQueue = dfsQueue[1:]

	for x := 1; x <= vertex; x++ {
		if arr[y][x] == 1 && dfsVisited[x] == 0 {
			dfsQueue = append(dfsQueue, x)
			_dfs()
		}
	}

}
