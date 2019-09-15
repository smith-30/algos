package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	n := nextLine()

	ok := true
	ss := strings.Split(n, "")

	for idx, item := range ss {
		if (idx+1)%2 != 0 {
			switch item {
			case "R", "U", "D":
			default:
				ok = false
				break
			}
		} else {
			switch item {
			case "L", "U", "D":
			default:
				ok = false
				break
			}
		}
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
