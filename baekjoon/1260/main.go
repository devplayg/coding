package main

import (
	"fmt"
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
