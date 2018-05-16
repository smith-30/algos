package column1

import (
	"fmt"
)

// 再帰は実行条件満たさなければ何もさせないということをして
// 次の処理に進めるようにする

// スライスを二分して再帰していくのに
// セパレータとして m を使い端点として活用している

// 再帰処理は内部で何回も呼ばれるので追うのが大変だけど
// 決まったルーチン(今回はmerge())を呼べるようにしてあげればいい

// エンドポイントというかメインの処理をきっちり決めて
// 内部のポインタに注意すれば他にも応用できそう

// MergeSort(target, left, m) が二分割した左辺の
// sort指示を出す、そこからは左辺が分解可能であれば呼ばれ続ける
// 後続のMergeSort(target, m+1, right)は右辺をsortし続ける

// 左の分解処理の最中でも、

// refs https://ide.geeksforgeeks.org/index.php
func MergeSort(target []int, left, right int) {
	if left < right {
		m := (left + (right - 1)) / 2
		fmt.Println("left: ", left, "right: ", right, "m: ", m)

		MergeSort(target, left, m)
		fmt.Println("left routine ended")
		MergeSort(target, m+1, right)
		fmt.Printf("%v, %v, %v\n", left, right, m)
		merge(target, left, m, right)
	}
}

func merge(target []int, l, m, r int) {
	n1 := m - l + 1
	n2 := r - m

	// create temp slice
	L := make([]int, n1, n1)
	R := make([]int, n2, n2)

	// Copy data to temp slice L[] and R[]
	for i := 0; i < n1; i++ {
		L[i] = target[l+i]
	}

	for j := 0; j < n2; j++ {
		R[j] = target[m+1+j]
	}

	fmt.Printf("L: %v\n", L)
	fmt.Printf("R: %v\n", R)

	// Merge the temp slice back into target[l..r]
	i := 0 // Initial index of first slice
	j := 0 // Initial index of second slice
	k := l // Initial index of merged slice

	for {
		ok := (i < n1 && j < n2)
		if ok {
			if L[i] <= R[j] {
				target[k] = L[i]
				i += 1
			} else {
				target[k] = R[j]
				j += 1
			}
			k += 1
		}

		if !ok {
			break
		}
	}

	fmt.Println(n1, n2)
	fmt.Println(i, j, k)
	fmt.Println()

	// Copy the remaining elements of L[], if there
	// are any
	for {
		if i < n1 {
			target[k] = L[i]
			i += 1
			k += 1
		} else {
			break
		}
	}

	// Copy the remaining elements of R[], if there
	// are any
	for {
		if j < n2 {
			target[k] = R[j]
			j += 1
			k += 1
		} else {
			break
		}
	}
}
