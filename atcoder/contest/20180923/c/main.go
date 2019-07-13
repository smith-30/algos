package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.ParseInt(sc.Text(), 10, 64)
	if e != nil {
		panic(e)
	}
	return int(i)
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	s, t := nextLine(), nextLine()

	n := len(s)

	var start, goal []int
	for i := 0; i < 26; i++ {
		start = append(start, -1)
		goal = append(goal, -1)
	}

	for i := 0; i < n; i++ {
		a := int(s[i] - 'a')
		b := int(t[i] - 'a')

		// 何かしら変換が行われている
		if start[a] != -1 || goal[b] != -1 {
			if start[a] != b || goal[b] != a {
				fmt.Println("No")
				return
			}
		} else {
			// 何も値がセットされていないとき
			start[a] = b
			goal[b] = a
		}

	}

	fmt.Println("Yes")
}
