package column2

// 再帰関数内であれこれ条件分岐あるのは追うのが大変なので多分よくない形なんだろうと思う
// 本来はあるスコープ内の範囲のことだけを繰り返しやらせて終了させればよいので
// 今回みたいに関数内で判定することが多くなるのは何か違う気がするな。。
// やっぱり、境界値データを最初に用意してそれが繰り返し処理として
// どうすればきれいに通るかを考えて上げたほうがよさそう

func DetectTaskA(src []int, numRange int, num *int) {
	if *num == 0 {
		separater := numRange / 2
		higher := make([]int, 0, len(src))
		lower := make([]int, 0, len(src))

		for _, val := range src {
			if val < separater {
				lower = append(lower, val)
				continue
			}
			if separater < val {
				higher = append(higher, val)
			}
		}

		// fmt.Println(lower)
		// fmt.Println(higher)
		// fmt.Println(separater)
		// fmt.Println()

		if len(higher) == 0 {
			*num = separater + 1
			DetectTaskA(higher, 0, num)
		}

		if len(lower) == 0 {
			*num = separater - 1
			DetectTaskA(higher, 0, num)
		}

		if len(lower) > len(higher) {
			if len(higher) == 1 {
				*num = higher[0] + 1
			}
			DetectTaskA(higher, numRange+separater, num)
		} else {
			if len(lower) == 1 {
				*num = lower[0] + 1
			}
			DetectTaskA(lower, separater, num)
		}
	}
}
