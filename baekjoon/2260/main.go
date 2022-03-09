package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var R, C int
var answer = -1
var dy = [4]int{-1, 0, 1, 0}
var dx = [4]int{0, 1, 0, -1}
var arr [][]int
var visited [][]bool
var broken [][]bool

//var breakDown [][]bool
var didBreakWall bool

type Pos struct {
	y, x, c int
}

func main() {
	fmt.Scanln(&R, &C)
	//fmt.Printf("r=%d, c=%d\n", R, C)
	arr = make([][]int, R+1)
	visited = make([][]bool, R+1)
	//breakDown = make([][]bool, R+1)
	for y := 1; y <= R; y++ {
		var line string
		fmt.Scanln(&line)
		list := toIntSlice(strings.SplitN(line, "", C))
		//spew.Dump(list)
		arr[y] = make([]int, C+1)
		visited[y] = make([]bool, C+1)
		//breakDown[y] = make([]bool, C+1)
		copy(arr[y][1:], list[:])
	}
	//spew.Dump(visited)

	dfs(Pos{1, 1, 1})
	fmt.Println(answer)

}

func dfs(pos Pos) {
	fmt.Printf("[%d-%d-%d]\n", pos.y, pos.x, pos.c)
	visited[pos.y][pos.x] = true
	defer func() {
		visited[pos.y][pos.x] = false
	}()

	if pos.y == R && pos.x == C {
		if pos.c == (R + C - 1) {
			fmt.Println("--------------- early")
			fmt.Println(answer)
			os.Exit(0)
		}
		//fmt.Printf("### %d\n", pos.c)
		if answer < 0 {
			answer = pos.c
			return
		}
		if pos.c < answer {
			answer = pos.c
		}
		return
	}

	for i := 0; i < 4; i++ {
		cy := pos.y + dy[i]
		cx := pos.x + dx[i]

		if cy >= 1 && cy <= R && cx >= 1 && cx <= C && !visited[cy][cx] {
			if arr[cy][cx] == 0 {
				dfs(Pos{cy, cx, pos.c + 1})
				continue
			}

			if !didBreakWall {
				//breakDown[cy][cx] = true
				arr[cy][cx] = 0
				didBreakWall = true
				//visited[cy][cx] = true
				dfs(Pos{cy, cx, pos.c + 1})
				arr[cy][cx] = 1
				didBreakWall = false
				//breakDown[cy][cx] = false
			}
		}
	}
}

func toIntSlice(strList []string) []int {
	intList := make([]int, len(strList))
	for i, c := range strList {
		intList[i], _ = strconv.Atoi(c)
	}
	return intList
}
