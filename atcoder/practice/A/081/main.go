package main

import (
	"fmt"
	"strings"
)

func main() {
	str := SingleStr()

	fmt.Println(strings.Count(str, "1"))
}

func SingleStr() string {
	var s string
	fmt.Scan(&s)
	return s
}
