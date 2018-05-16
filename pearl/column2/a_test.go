package column2

import (
	"testing"
)

func TestDetectTaskA(t *testing.T) {
	type args struct {
		src      []int
		numRange int
		num      int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ok_1",
			args: args{
				src:      []int{1, 2, 3, 7, 8, 9, 10},
				numRange: 15,
				num:      0,
			},
		},
		{
			name: "ok_2",
			args: args{
				src:      []int{1, 3, 7, 8, 9, 10},
				numRange: 15,
				num:      0,
			},
		},
		{
			name: "ok_3",
			args: args{
				src:      []int{1, 2, 3, 4, 5, 7, 8, 9, 10},
				numRange: 15,
				num:      0,
			},
		},
		{
			name: "ok_4",
			args: args{
				src:      []int{1, 2},
				numRange: 15,
				num:      0,
			},
		},
		{
			name: "ok_5",
			args: args{
				src:      []int{7, 8, 9},
				numRange: 15,
				num:      0,
			},
		},
		{
			name: "ok_6",
			args: args{
				src:      []int{1, 9, 12, 15, 18, 5, 3, 2, 16, 11, 10, 7, 17},
				numRange: 20,
				num:      0,
			},
		},
	}
	for _, tt := range tests {
		r := false
		DetectTaskA(tt.args.src, tt.args.numRange, &tt.args.num)
		for _, val := range tt.args.src {
			if val == tt.args.num {
				r = true
			}
		}
		if r {
			t.Errorf(`
				name: %v
				failed. act. %v
			`,
				tt.name,
				tt.args.num)
		}

	}
}
