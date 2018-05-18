package plg

import (
	"fmt"
	"sort"
	"strings"
)

var (
	dummyDict = []string{"from", "time", "item", "form", "off", "test"}
)

func Anagram() {
	// 容量割り当てときたいが目安がわからないので保留
	list := make(map[string][]string)

	for _, word := range dummyDict {
		key := sortStr(word)
		list[key] = append(list[key], word)
	}

	for _, words := range list {
		for _, w := range words {
			fmt.Print(w, " ")
		}
		fmt.Println()
	}
}

func sortStr(k string) string {
	s := strings.Split(k, "")
	sort.Strings(s)

	return strings.Join(s, "")
}
