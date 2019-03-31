package main

import (
	"fmt"
	"strings"
)

func main() {
	h := SingleInt()
	w := SingleInt()

	points := map[int]map[int]string{}
	for i := 0; i < h; i++ {
		points[i] = map[int]string{}
		str := strings.Split(SingleStr(), "")
		for idx, item := range str {
			points[i][idx] = item
		}
	}

	rowRange := []int{-1, 0, 1}
	lineRange := []int{-1, 0, 1}

	for i := 0; i < h; i++ {
		var line string
		for j := 0; j < w; j++ {
			if points[i][j] == "#" {
				line += points[i][j]
				continue
			}
			var p int
			for _, r := range rowRange {
				for _, l := range lineRange {
					if points[i+r][j+l] == "#" {
						p++
					}
				}
			}
			if p == 0 {
				line += "0"
			} else {
				line += fmt.Sprintf("%d", p)
			}

		}
		fmt.Printf("%v\n", line)
	}

}

func SingleInt() int {
	var n int
	fmt.Scan(&n)
	return n
}
func SingleStr() string {
	var s string
	fmt.Scan(&s)
	return s
}

func ScanStrings(len int) (strings []string) {
	var str string
	for i := 0; i < len; i++ {
		fmt.Scanf("%s", &str)
		strings = append(strings, str)
	}
	return
}
