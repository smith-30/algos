package counting_sort

func counting_sort(x []int, k int) []int {
	result := make([]int, len(x))
	countMap := make([]int, k)

	// create map (O(n))
	for _, val := range x {
		countMap[val]++
	}

	// change countMap's value.
	var prev int
	for k, val := range countMap {
		prev += val
		countMap[k] = prev
	}

	// while decreasing a, put the value of x in result. (O(n))
	for _, val := range x {
		countMap[val]--
		result[countMap[val]] = val
	}

	return result
}
