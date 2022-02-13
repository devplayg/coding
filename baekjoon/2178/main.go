package main

import (
	"fmt"
)

type Pos struct {
	y int
	x int
	c int
}

var N int
var M int
var arr = [101][101]int{}
var visited = [101][101]bool{}
var queue = make([]*Pos, 0)
var dy = [4]int{1, -1, 0, 0}
var dx = [4]int{0, 0, 1, -1}

func valid(y, x int) bool {
	if !(y > 0 && y <= N && x > 0 && x <= M) {
		return false
	}
	if arr[y][x] == 0 {
		return false
	}

	if visited[y][x] {
		return false
	}
	return true
}

func main() {
	fmt.Scanf("%d %d", &N, &M)
	for y := 1; y <= N; y++ {
		for x := 1; x <= M; x++ {
			if x == 1 {
				fmt.Scanf("\n%1d", &arr[y][x])
				continue
			}
			fmt.Scanf("%1d", &arr[y][x])
		}
	}

	queue = append(queue, &Pos{1, 1, 1})
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		if pos.y == N && pos.x == M {
			fmt.Println(pos.c)
			return
		}
		c := pos.c + 1
		for i := 0; i < 4; i++ {
			if valid(pos.y+dy[i], pos.x+dx[i]) {
				queue = append(queue, &Pos{pos.y + dy[i], pos.x + dx[i], c})
				visited[pos.y+dy[i]][pos.x+dx[i]] = true
			}
		}
	}
}
