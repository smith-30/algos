package main

import "fmt"

func main() {
	var n, a, b, c int
	fmt.Scan(&n)
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	cand := []int{a, b, c}

	list := map[int]int{}
	for index := 0; index < n; index++ {
		var l int
		fmt.Scan(&l)

		list[l] = l
	}

	cand2 := []int{}
	for _, item := range cand {
		if _, ok := list[item]; ok {
			delete(list, item)
		} else {
			cand2 = append(cand2, item)
		}
	}

	var re



	fmt.Printf("%#v\n", cand2)

	var re int
	fmt.Println(re)
}
