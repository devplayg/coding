package main

import (
	"fmt"
	"sort"
)

func main() {
	// Input
	var nodeCount, lineCount, startId int
	fmt.Scanf("%d %d %d", &nodeCount, &lineCount, &startId)

	// Create nodes and draw lines
	myNodeMap := newNodeMap(nodeCount)

	var line1, line2 int
	for i := 0; i < lineCount; i++ {
		fmt.Scanf("%d %d", &line1, &line2)
		myNodeMap.m[line1].link(myNodeMap.m[line2])
	}

	myNodeMap.dfs(startId)

	fmt.Println()
	myNodeMap.clear()

	myNodeMap.bfs(startId)
}

// node
type node struct {
	id       int
	children []*node
	visited  bool
}

func (t *node) link(a *node) {
	t.children = append(t.children, a)
	a.children = append(a.children, t)
}

// nodeMap
type nodeMap struct {
	m     map[int]*node
	queue []*node
	idx   int
}

func newNodeMap(count int) *nodeMap {
	m := make(map[int]*node)
	for n := 1; n <= count; n++ {
		node := &node{
			id:       n,
			children: make([]*node, 0),
		}
		m[n] = node
	}
	return &nodeMap{m: m, queue: make([]*node, 0)}
}

func (nm *nodeMap) clear() {
	for k, _ := range nm.m {
		nm.m[k].visited = false
	}
}

func (nm *nodeMap) dfs(id int) {
	if nm.m[id].visited {
		return
	}

	node := nm.m[id]
	fmt.Printf("%d ", node.id)
	node.visited = true

	// Sort
	sort.Slice(node.children, func(i, j int) bool {
		return node.children[i].id < node.children[j].id
	})

	for i := 0; i < len(node.children); i++ {
		nm.dfs(node.children[i].id)
	}
}

func (nm *nodeMap) bfs(id int) {
	nm.visitByBfs(id)
	for nm.idx < len(nm.queue)-1 || len(nm.queue) == 1 {
		nm.visitByBfs(nm.queue[nm.idx].id)
		nm.idx++
	}
}

func (nm *nodeMap) visitByBfs(id int) {
	node := nm.m[id]
	if node.visited {
		return
	}
	fmt.Printf("%d ", node.id)
	node.visited = true

	for i := 0; i < len(node.children); i++ {
		nm.queue = append(nm.queue, node.children[i])
	}
}
