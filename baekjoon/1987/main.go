package main

import (
	"fmt"
	"strings"
)

const MAX = 20

var arr [MAX + 1][MAX + 1]string
var maxNodeCount = 0

var rc, cc int
var dy = [4]int{-1, 0, 1, 0}
var dx = [4]int{0, 1, 0, -1}

type node struct {
	y       int
	x       int
	count   int
	parent  *node
	visited map[string]bool
}

func newNode(y, x, count int, parent *node) *node {
	if parent == nil {
		return &node{
			y,
			x,
			count,
			parent,
			map[string]bool{
				arr[y][x]: true,
			},
		}
	}
	visited := copyMap(parent.visited)
	visited[arr[y][x]] = true
	return &node{
		y,
		x,
		count,
		parent,
		visited,
	}

}

func (n *node) char() string {
	return arr[n.y][n.x]
}

func (n *node) print() {
	fmt.Printf("  - [%d-%d-%d] %s\n", n.y, n.x, n.count, n.char())
}

func (n *node) done() {
	if maxNodeCount < n.count {
		maxNodeCount = n.count
	}
}

var queue = make([]*node, 0)
var root *node

func main() {
	fmt.Scanln(&rc, &cc)
	for y := 1; y <= rc; y++ {
		var line string
		fmt.Scanln(&line)
		colArr := strings.SplitN(line, "", cc)
		copy(arr[y][1:], colArr[0:cc])
	}

	root = newNode(1, 1, 1, nil)
	queue = append(queue, root)
	traverse()
	fmt.Println(maxNodeCount)
}

func copyMap(m map[string]bool) map[string]bool {
	cp := make(map[string]bool)
	for k, v := range m {
		cp[k] = v
	}
	return cp
}

func traverse() {
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		for i := 0; i < 4; i++ {
			ny, nx := n.y+dy[i], n.x+dx[i]

			if valid(ny, nx) {
				if hasFootPrints(n, arr[ny][nx]) {
					n.done()
				} else {
					//newNode := &node{ny, nx, n.count + 1, n, make(map[string]int)}
					newNode := newNode(ny, nx, n.count+1, n)
					queue = append(queue, newNode)
				}
			}
		}
	}
}

func hasFootPrints(parent *node, char string) bool {
	if _, has := parent.visited[char]; has {
		return true
	}
	return false
}

func valid(y, x int) bool {
	if y < 1 || y > MAX {
		return false
	}
	if x < 1 || x > MAX {
		return false
	}
	if arr[y][x] == "" {
		return false
	}
	return true
}
