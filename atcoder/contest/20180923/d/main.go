package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

const MOD = 1000000007

func main() {
	N := int64(nextInt())
	M := int64(nextInt())

	MNokori := int(M)
	var ans int64 = 1
	
	// 素因数分解
	for i := 2; i*i <= MNokori; i++ {
		if MNokori%i == 0 {
			var cnt int64
			for MNokori%i == 0 {
				cnt++
				MNokori /= i
			}
			//cntが2^Xとか3^XのXの部分
			ans *= calcComb(cnt+N-1, N-1) //N-1はcntと同じ
			ans %= MOD
		}
	}
	if MNokori != 1 {
		//最後に素数が残ってる分を処理する
		ans *= calcComb(N+1-1, N-1) //N-1はcntと同じ
		ans %= MOD
	}

	fmt.Printf("%#v\n", ans)
}

func modpow(a, p int64) int64 {
	if p == 0 {
		return 1
	}

	if p%2 == 0 {
		//pが偶数の時
		halfP := p / 2
		half := modpow(a, halfP)
		//a^(p/2) をhalfとして、half*halfを計算
		return half * half % MOD
	} else {
		//pが奇数の時は、
		//pを偶数にするために1減らす
		return a * modpow(a, p-1) % MOD
	}
}

//(10*9*8)/(3*2*1);
//10*9*8 -> ansMul
//3*2*1 -> ansDiv
func calcComb(a, b int64) int64 {
	if b > a-b {
		return calcComb(a, a-b)
	}

	var ansMul int64 = 1
	var ansDiv int64 = 1
	var i int64
	for i = 0; i < b; i++ {
		ansMul *= (a - i)
		ansDiv *= (i + 1)
		ansMul %= MOD
		ansDiv %= MOD
	}
	//ansMul / ansDivをやりたい
	//ansDivの逆元を使って求めよう！

	ans := ansMul * modpow(ansDiv, MOD-2) % MOD
	return ans
}
