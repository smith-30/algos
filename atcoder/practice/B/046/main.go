package main

import (
	"fmt"
	"math"
)

func main() {
	n, k := SingleInt(), SingleInt()

	// まず左端のボールに塗る色を決めます。これは K 通りあります。さらに左から順番に塗る色を決めていき
	// ます。この時、これまでどのように塗っていようが、”ひとつ左のボールに塗られている色”以外の K − 1 通
	// りの色が塗れるので、全体では K · (K − 1)^N−1 通りになります

	// 一番左の選び方K通り、その隣の選び方は、K-1通り。またその隣の塗り方はK-1通り
	// K * (K-1) * (K-1) ... (K-1)は(N-1)個現れる
	// (K-1)*(N-1)　ではないのは、順列であるため

	fmt.Println(k * int(math.Pow(float64(k-1), float64(n-1))))
}

func SingleInt() int {
	var n int
	fmt.Scan(&n)
	return n
}
