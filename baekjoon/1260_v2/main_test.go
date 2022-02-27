package main

//
//func setArr(from, to int) {
//	arr[from][to] = 1
//	arr[to][from] = 1
//}
//
//func TestCases(t *testing.T) {
//	tt := []struct {
//		description string
//		vertex, edge, start        int
//		expect      string
//	}{
//		{"Empty", 1000,1,1000, "1000 999 "},
//	}
//
//	for _, tc := range tt {
//		t.Run(tc.description, func(t *testing.T) {
//			arr = [MAX][MAX]int{}
//			bfsVisited = [MAX]int{}
//			dfsVisited = [MAX]int{}
//
//			output := bytes.Buffer{}
//			w = &output
//
//			setArr(999, 1000)
//			dfs(tc.start)
//			expected := "1000 999 "
//			if output.String() != expected {
//				t.Errorf("expected=%s, actual=%s", expected, output.String())
//			}
//
//			//var output bytes.Buffer
//			//SayHello(&output, tc.name)
//			//if tc.expect != output.String() {
//			//	t.Errorf("got %s but expected %s", output.String(), tc.expect)
//			//}
//		})
//	}
//
//}
//
//func TestAlpha(t *testing.T) {
//	 arr = [MAX][MAX]int{}
//	 bfsVisited = [MAX]int{}
//	 dfsVisited = [MAX]int{}
//
//	output := bytes.Buffer{}
//	w = &output
//
//	vertex, edge, start = 1000, 1, 1000
//	setArr(999, 1000)
//	dfs(1000)
//	expected := "1000 999 "
//	if output.String() != expected {
//		t.Errorf("expected=%s, actual=%s", expected, output.String())
//	}
//
//}
