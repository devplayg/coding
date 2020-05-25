/*
 https://www.acmicpc.net/problem/1978
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input int
	fmt.Scanf("%d", &input)

	if input < 0 || input > 100 {
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	var line string
	if scanner.Scan() {
		line = scanner.Text()
	}
	nums := strings.SplitN(line, " ", input)
	if len(nums) != input {
		return
	}

	var count int
	for _, str := range nums {
		n, err := strconv.Atoi(str)
		//if n  < 0 || n > 1000 {
		//	return
		//}
		if err != nil {
			return
		}
		if isPrimeNumber(n) {
			count++
		}
	}
	fmt.Println(count)
}

func isPrimeNumber(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			if i != n {
				return false
			}
			return true
		}
	}
	return false
}
