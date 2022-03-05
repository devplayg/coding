package main

import (
	"fmt"
	"strings"
)

var R, C int
var visited = [26]bool{}
var arr [][]int
var dy = [4]int{-1, 0, 1, 0}
var dx = [4]int{0, 1, 0, -1}
var answer = 0

func main() {
	fmt.Scanln(&R, &C)
	arr = make([][]int, R)
	for y := 0; y < R; y++ {
		var line string
		fmt.Scanln(&line)
		colArr := strings.SplitN(line, "", C)
		arr[y] = toIntSlice(colArr)
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

func toIntSlice(strList []string) []int {
	intList := make([]int, len(strList))
	for i, c := range strList {
		intList[i] = int(c[0]) - int('A')
	}
	return intList
}
