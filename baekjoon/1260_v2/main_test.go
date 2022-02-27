package main

import (
	"testing"
)

func setArr(from, to int) {
	arr[from][to] = 1
	arr[to][from] = 1
}

func TestAlpha(t *testing.T) {
	vertex, edge, start = 1000, 1, 1000
	setArr(999, 1000)
	result := dfs(1000)

	expected := "1000 999 "
	if result != expected {
		t.Errorf("expected=%s, actual=%s", expected, result)
	}

}
