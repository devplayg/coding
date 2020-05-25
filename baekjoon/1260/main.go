package main

import (
	"fmt"
	"sort"
)

type node struct {
	val      int
	visited  bool
	children []*node
}

func newNode(val int) *node {
	return &node{
		val:      val,
		children: nil,
	}
}

func (t *node) append(a *node) {
	if t.children == nil {
		t.children = make([]*node, 0)
	}
	t.children = append(t.children, a)

	if a.children == nil {
		a.children = make([]*node, 0)
	}
	a.children = append(a.children, t)
}

type nodeMap map[int]*node

func (nm nodeMap) clear() {
	for k, _ := range nm {
		nm[k].visited = false
	}

}

func main() {
	// Input
	var c1, c2, c3 int
	fmt.Scanf("%d %d %d", &c1, &c2, &c3)

	// Create nodes and draw lines
	myNodeMap := createNodes(c1)

	var line1, line2 int
	for i := 0; i < c2; i++ {
		fmt.Scanf("%d %d", &line1, &line2)
		myNodeMap[line1].append(myNodeMap[line2])
	}

	visitByDFS(myNodeMap[c3])
	fmt.Println()

	myNodeMap.clear()
	searchByBFS(myNodeMap[c3])
	fmt.Println()
}

func createNodes(count int) nodeMap {
	m := make(nodeMap)
	for i := 1; i <= count; i++ {
		node := newNode(i)
		m[i] = node
	}
	return m
}

func visitByDFS(n *node) {
	if n.visited {
		return
	}

	fmt.Printf("%d ", n.val)
	n.visited = true

	if n.children == nil {
		return
	}

	sort.Slice(n.children, func(i, j int) bool {
		return n.children[i].val < n.children[j].val
	})

	for i := 0; i < len(n.children); i++ {
		visitByDFS(n.children[i])
	}
}

var queue = make([]*node, 0)
var idx = 0

func searchByBFS(n *node) {
	visitByBFS(n)
	for idx < len(queue)-1 || len(queue) == 1 {
		visitByBFS(queue[idx])
		idx++
	}
}
func visitByBFS(n *node) {
	if n.visited {
		return
	}

	fmt.Printf("%d ", n.val)
	n.visited = true

	if n.children == nil {
		return
	}

	for i := 0; i < len(n.children); i++ {
		queue = append(queue, n.children[i])
	}
}
