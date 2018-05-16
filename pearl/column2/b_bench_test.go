package column2

import (
	"fmt"
	"testing"

	"github.com/smith-30/algo/pearl/bench_helper"
)

var (
	refIdx     = map[int][]int{}
	dataCounts = []int{100, 200, 400, 800, 1600, 3200, 6400, 12800, 25600, 51200}
)

func init() {
	for _, dataCount := range dataCounts {
		refIdx[dataCount] = []int{1, dataCount / 2, dataCount - 1}
	}
}

func Benchmark_RotateTaskB(b *testing.B) {
	for _, dataCount := range dataCounts {
		for _, idx := range refIdx[dataCount] {
			b.Run(fmt.Sprintf("dataCount: %v, idx: %v", dataCount, idx), func(b *testing.B) {
				s := bench_helper.GenStrSlice(dataCount)
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					RotateTaskB(s, idx)
				}
			})
		}
		fmt.Println()
	}
}

func Benchmark_NewRotateTaskB(b *testing.B) {
	for _, dataCount := range dataCounts {
		for _, idx := range refIdx[dataCount] {
			b.Run(fmt.Sprintf("dataCount: %v, idx: %v", dataCount, idx), func(b *testing.B) {
				s := bench_helper.GenStrSlice(dataCount)
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					NewRotateTaskB(s, idx)
				}
			})
		}
		fmt.Println()
	}
}
