package main

import "fmt"

func main() {
	n := SingleInt()
	var s string
	fmt.Scan(&s)


	var white, black int
	whiteArr := make([]int, n)
	blackArr := make([]int, n)

	//
	// 左からn個を白、以降は黒という風に分けることを目指す
	// そうすると、最初のn個は、黒の数を数え、それ以降は白の数を
	// 数えていけば、白黒の境目がiのときに入れ替える数がわかる
	//
	for i := 0; i < n; i++ {
		// right hand side boundary
		if "." == string(s[i]) {
			white++
		}
		// left hand side boundary
		if "#" == string(s[i]) {
			black++
		}
		whiteArr[i] = white
		blackArr[i] = black
	}
	var result int
	for i := 0; i < n; i++ {
		boundary := blackArr[i] + whiteArr[n-1] - whiteArr[i]
		fmt.Printf("boundary %#v\n", boundary)
		if i == 0 {
			result = boundary
		}
		if result > boundary {
			result = boundary
		}
	}
	if result > whiteArr[n-1] {
		fmt.Println(whiteArr[n-1])
		return
	}
	fmt.Printf("%#v", "*")
	fmt.Println(result)
}

func SingleInt() int {
	var n int
	fmt.Scan(&n)
	return n
}
