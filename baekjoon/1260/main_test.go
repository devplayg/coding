package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	type args struct {
		nodeCount int
		lineCount int
		startNo   int
		lines     [][2]int
		want      [2]string
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "problem-1", args: struct {
			nodeCount int
			lineCount int
			startNo   int
			lines     [][2]int
			want      [2]string
		}{nodeCount: 4, lineCount: 5, startNo: 1, lines: [][2]int{
			{1, 2},
			{1, 3},
			{1, 4},
			{2, 4},
			{3, 4},
		}, want: [2]string{"1 2 4 3", "1 2 3 4"}}},

		{name: "problem-2", args: struct {
			nodeCount int
			lineCount int
			startNo   int
			lines     [][2]int
			want      [2]string
		}{nodeCount: 5, lineCount: 5, startNo: 3, lines: [][2]int{
			{5, 4},
			{5, 2},
			{1, 2},
			{3, 4},
			{3, 1},
		}, want: [2]string{"3 1 2 5 4", "3 1 4 2 5"}}},

		{name: "problem-3", args: struct {
			nodeCount int
			lineCount int
			startNo   int
			lines     [][2]int
			want      [2]string
		}{nodeCount: 1000, lineCount: 1, startNo: 1000, lines: [][2]int{
			{999, 1000},
		}, want: [2]string{"1000 999", "1000 999"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue = make([]*node, 0)
			idx = 0
			m := createNodes(tt.args.nodeCount)
			for _, line := range tt.args.lines {
				m[line[0]].append(m[line[1]])
			}
			startCaputringStdout()
			visitByDFS(m[tt.args.startNo])
			res := stopCaptureStdout()
			if strings.TrimSpace(res) != tt.args.want[0] {
				t.Errorf("(DFS) got = %s, want = %s", res, tt.args.want[0])
			}

			m.clear()

			startCaputringStdout()
			searchByBFS(m[tt.args.startNo])
			res = stopCaptureStdout()
			if strings.TrimSpace(res) != tt.args.want[1] {
				t.Errorf("(BFS) got = %s, want = %s", res, tt.args.want[1])
			}
		})
	}

}

var r, w, oldStdout *os.File

func startCaputringStdout() {
	oldStdout = os.Stdout
	r, w, _ = os.Pipe()
	os.Stdout = w
}

func stopCaptureStdout() string {
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = oldStdout
	return string(bytes.TrimSpace(out))
}
