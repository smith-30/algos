package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var re float64

	for index := 0; index < n; index++ {
		var val float64
		var coin string

		fmt.Scan(&val)
		fmt.Scan(&coin)

		switch coin {
		case "JPY":
			re += val
		case "BTC":
			re += val * 380000
		}

	}
	fmt.Printf("%.9f\n", re)
}
