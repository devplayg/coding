package main

import (
	"fmt"
	"sort"
	"strings"
)

const MAX = 25 + 1

var arr [MAX][MAX]string
var visited [MAX][MAX]bool
var size int
var dy = [4]int{-1, 0, 1, 0}
var dx = [4]int{0, 1, 0, -1}

type Pos struct {
	y int
	x int
}

func main() {
	fmt.Scanln(&size)
	for y := 1; y <= size; y++ {
		var line string
		fmt.Scanln(&line)
		cols := strings.SplitN(line, "", size)
		copy(arr[y][1:], cols[0:size])
	}
	//spew.Dump(arr)

	gid := 0
	list := make([]int, 0)
	for y := 1; y <= size; y++ {
		for x := 1; x <= size; x++ {
			if arr[y][x] == "1" && !visited[y][x] {
				queue = make([]Pos, 0)
				gid++
				count := findGroup(Pos{y, x})
				list = append(list, count)
			}
		}
	}

	sort.Ints(list)
	fmt.Println(len(list))
	for _, c := range list {
		fmt.Println(c)
	}
}

var queue = make([]Pos, 0)

func valid(y, x int) bool {
	if y < 1 || y >= MAX {
		return false
	}
	if x < 1 || x >= MAX {
		return false
	}
	if visited[y][x] {
		return false
	}
	if arr[y][x] == "1" {
		return true
	}
	return false
}

func findGroup(pos Pos) int {
	queue = append(queue, pos)
	visited[pos.y][pos.x] = true
	//fmt.Printf("[%d,%d] visited\n", pos.y, pos.x)
	count := 0

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		count++
		for i := 0; i < 4; i++ {
			if valid(p.y+dy[i], p.x+dx[i]) {
				np := Pos{p.y + dy[i], p.x + dx[i]}
				queue = append(queue, np)
				visited[np.y][np.x] = true
				//fmt.Printf("[%d,%d] visited\n", np.y, np.x)
			}
		}
	}
	return count
}
