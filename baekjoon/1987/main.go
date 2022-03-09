package main

import (
	"fmt"
)

var R, C int
var visited = [26]bool{}
var arr [][]int
var dy = [4]int{-1, 0, 1, 0}
var dx = [4]int{0, 1, 0, -1}
var answer = 0
var s string

type newline struct{ tok string }

func main() {
	fmt.Scanln(&R, &C)
	arr = make([][]int, R)
	for y := 0; y < R; y++ {
		arr[y] = make([]int, C)
		for x := 0; x < C; x++ {
			fmt.Scanf("%1s", &s)
			arr[y][x] = int(s[0]) - int('A')
		}
		fmt.Scanf("%1s")
	}
	dfs(0, 0, 0)
	fmt.Printf("%d\n", answer)
}

func dfs(y, x, c int) {
	//fmt.Printf("%d-%d-%d-%d\n", y, x, c, arr[y][x])
	if visited[arr[y][x]] {
		if c > answer {
			answer = c
		}
		return
	}

	visited[arr[y][x]] = true
	for i := 0; i < 4; i++ {
		cy := y + dy[i]
		cx := x + dx[i]

		if cy >= 0 && cx >= 0 && cy < R && cx < C {
			dfs(cy, cx, c+1)
		}
	}
	visited[arr[y][x]] = false
}
