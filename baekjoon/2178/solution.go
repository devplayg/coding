package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"strconv"
	"strings"
)

type Maze struct {
	width  int
	height int
	arr    [][]int
}

func NewMaze(h, w int, mazeStr string) *Maze {
	arr := make([][]int, h)
	for i := range arr {
		arr[i] = make([]int, w)
	}
	//spew.Dump(mazeStr)
	rows := strings.Split(mazeStr, "\n")
	spew.Dump(rows)
	for i := range rows {
		for j := range rows[i] {
			n, _ := strconv.Atoi(string(rows[i][j]))
			if n == 1 {
				arr[i][j] = 1
			}
		}
		//fmt.Println()
	}

	maze := Maze{
		width:  0,
		height: 0,
		arr:    arr,
	}

	return &maze
}

func (m *Maze) display() {
	fmt.Println("==================")
	for i := range m.arr {
		for j := range m.arr[i] {
			fmt.Printf("%d", m.arr[i][j])
		}
		fmt.Println()
	}
}

type Cell struct {
}

func (c *Cell) hasTop() bool {
	return true
}

func (c *Cell) hasRight() bool {
	return true

}

func (c *Cell) hasBottom() bool {
	return true
}

func (c *Cell) hasLeft() bool {
	return true
}
