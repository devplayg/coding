package main

import (
	"fmt"
	"sort"
)

var queue []*node

// Using map
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

	dfs(myNodeMap.m[startId])

	fmt.Println()
	myNodeMap.clear()

	myNodeMap.bfs(myNodeMap.m[startId])
}

// Using slice
//func main() {
//	// Input
//	var nodeCount, lineCount, startId int
//	fmt.Scanf("%d %d %d", &nodeCount, &lineCount, &startId)
//
//	// Create nodes and draw lines
//	nl := newNodeList(nodeCount)
//
//	var line1, line2 int
//	for i := 0; i < lineCount; i++ {
//		fmt.Scanf("%d %d", &line1, &line2)
//		nl.list[line1-1].link(nl.list[line2-1])
//	}
//
//	dfs(nl.list[startId-1])
//
//	fmt.Println()
//	nl.clear()
//
//	nl.bfs(nl.list[startId-1])
//}

// node
type node struct {
	id       int
	children []*node
	visited  bool
}

func (n *node) link(a *node) {
	n.children = append(n.children, a)
	a.children = append(a.children, n)
}

func (n *node) visit() {
	fmt.Printf("%d ", n.id)
	n.visited = true
}

// nodeMap
type nodeMap struct {
	m   map[int]*node
	idx int
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
	return &nodeMap{m: m}
}

func (nm *nodeMap) clear() {
	for k, _ := range nm.m {
		nm.m[k].visited = false
	}
	queue = make([]*node, 0)
}

func (nm *nodeMap) bfs(nd *node) {
	nm.visitByBfs(nd)
	for nm.idx < len(queue)-1 || len(queue) == 1 {
		nm.visitByBfs(queue[nm.idx])
		nm.idx++
	}
}

func (nm *nodeMap) visitByBfs(nd *node) {
	if nd.visited {
		return
	}
	nd.visit()

	for i := 0; i < len(nd.children); i++ {
		queue = append(queue, nd.children[i])
	}
}

// nodeList
type nodeList struct {
	list  []*node
	queue []*node
	idx   int
}

func newNodeList(count int) *nodeList {
	nl := nodeList{
		list:  make([]*node, count),
		queue: make([]*node, 0),
		idx:   0,
	}

	for i := 0; i < count; i++ {
		nl.list[i] = &node{
			id:       i + 1,
			children: make([]*node, 0),
			visited:  false,
		}
	}

	return &nl
}

func (nl *nodeList) clear() {
	for i := range nl.list {
		nl.list[i].visited = false
	}
	queue = make([]*node, 0)
}

func (nl *nodeList) bfs(nd *node) {
	nl.visitByBfs(nd)

	for i := 0; i < len(nd.children); i++ {
		nl.visitByBfs(nd.children[i])
	}

	for i := 0; i < len(nl.queue); i++ {
		nl.visitByBfs(nl.queue[i])
	}
}

func (nl *nodeList) visitByBfs(nd *node) {
	if nd.visited {
		return
	}

	nd.visit()
	for i := range nd.children {
		nl.queue = append(nl.queue, nd.children[i])
	}
}

func dfs(nd *node) {
	if nd.visited {
		return
	}

	nd.visit()

	if len(nd.children) < 1 {
		return
	}

	// Sort
	sort.Slice(nd.children, func(i, j int) bool {
		return nd.children[i].id < nd.children[j].id
	})

	for i := 0; i < len(nd.children); i++ {
		dfs(nd.children[i])
	}
}
