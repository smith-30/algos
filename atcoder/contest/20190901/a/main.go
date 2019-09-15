package main

import (
	"fmt"
	"strings"
)

// var sc = bufio.NewScanner(os.Stdin)

// func init() {
// 	sc.Split(bufio.ScanWords)
// }

// func nextInt() int {
// 	sc.Scan()
// 	i, e := strconv.ParseInt(sc.Text(), 10, 64)
// 	if e != nil {
// 		panic(e)
// 	}
// 	return int(i)
// }

func SingleStr() string {
	var s string
	fmt.Scan(&s)
	return s
}

func main() {
	s := SingleStr()
	ss := strings.Split(s, "")
	n := SingleStr()
	nn := strings.Split(n, "")

	var re int

	for index := 0; index < 3; index++ {
		if ss[index] == nn[index] {
			re++
		}
	}
	fmt.Println(re)
}
