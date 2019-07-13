package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var (
		s string
		t string
	)
	fmt.Scan(&s, &t)

	sMap := make(map[rune]int)
	tMap := make(map[rune]int)
	for _, r := range s {
		sMap[r]++
	}
	for _, r := range t {
		tMap[r]++
	}

	var (
		sVals []int
		tVals []int
	)

	for _, v := range sMap {
		sVals = append(sVals, v)
	}
	for _, v := range tMap {
		tVals = append(tVals, v)
	}

	sort.Ints(sVals)
	sort.Ints(tVals)

	for i := 0; i < len(sVals); i++ {
		if sVals[i] != tVals[i] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
