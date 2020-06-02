package main

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text())
	//}
	//
	//if scanner.Err() != nil {
	//	// handle error.
	//}
	//
	//fmt.Scan
	//4 6
	//101111
	//101010
	//101011
	//111011
	input := `
101111
101010
101011
111011
`
	m := NewMaze(4, 6, input)
	m.display()
	//spew.Dump(m)
}
