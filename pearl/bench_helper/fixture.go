package bench_helper

func GenStrSlice(count int) []string {
	s := make([]string, 0, count)

	for i := 0; i < count; i++ {
		s = append(s, "a")
	}
	return s
}

func GenIntSlice(count int) []int {
	ii := make([]int, 0, count)

	for i := 0; i < count; i++ {
		ii = append(ii, 1)
	}
	return ii
}
