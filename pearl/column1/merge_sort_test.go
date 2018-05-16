package column1

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		target []int
		left   int
		right  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ok_6",
			args: args{
				target: []int{12, 11, 13, 5},
				left:   0,
				right:  3,
			},
			want: []int{5, 11, 12, 13},
		},
		{
			name: "ok_7",
			args: args{
				target: []int{38, 27, 43, 3, 9, 82, 10},
				left:   0,
				right:  6,
			},
			want: []int{3, 9, 10, 27, 38, 43, 82},
		},
		{
			name: "ok_2",
			args: args{
				target: []int{1},
				left:   0,
				right:  0,
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.args.target, tt.args.left, tt.args.right)
			if !reflect.DeepEqual(tt.args.target, tt.want) {
				t.Errorf(`
				failed.
				act. %v
				exp. %v
			`, tt.args.target, tt.want)
			}
		})
	}
}
