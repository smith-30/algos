package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	var sum int
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		sum += a[i]
	}

	mod := 998244353

	combinations := make([][]int, n)
	combinations[0] = make([]int, sum+1)
	combinations[0][0] = 2
	combinations[0][a[0]] = 1
	for i := 0; i < n-1; i++ {
		combinations[i+1] = make([]int, sum+1)
		for j := 0; j <= sum; j++ {
			combinations[i+1][j] = combinations[i][j] * 2 % mod
			if j-a[i+1] >= 0 {
				combinations[i+1][j] += combinations[i][j-a[i+1]] % mod
			}
		}
	}

	for idx, item := range combinations {
		fmt.Printf("%v, %#v\n", idx, item)
	}

	dups := make([][]int, n)
	dups[0] = make([]int, sum+1)
	dups[0][0] = 1
	dups[0][a[0]] = 1
	for i := 0; i < n-1; i++ {
		dups[i+1] = make([]int, sum+1)
		for j := 0; j <= sum; j++ {
			dups[i+1][j] = dups[i][j] % mod
			if j-a[i+1] >= 0 {
				dups[i+1][j] += dups[i][j-a[i+1]] % mod
			}
		}
	}

	bigThree := big.NewInt(3)
	bigNum := big.NewInt(int64(n))
	totalCombinations := bigThree.Exp(bigThree, bigNum, nil)
	half := sum / 2
	if sum%2 != 0 {
		half++
	}
	for i := half; i <= sum; i++ {
		fmt.Printf("comb:::%#v\n", int64(combinations[n-1][i]*3))
		totalCombinations = totalCombinations.Sub(totalCombinations, big.NewInt(int64(combinations[n-1][i]*3)))
		if i == sum/2 {
			totalCombinations = totalCombinations.Add(totalCombinations, big.NewInt(int64(dups[n-1][i]*3)))
		}
	}

	ret := totalCombinations.Mod(totalCombinations, big.NewInt(int64(998244353)))
	fmt.Println(ret)
}
