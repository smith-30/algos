package counting_sort

import (
	"reflect"
	"testing"
)

func Test_counting_sort(t *testing.T) {
	type args struct {
		x []int
		k int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ok_0",
			args: args{
				x: []int{3, 4, 2, 1, 0, 0, 4, 3, 4, 2},
				k: 5,
			},
			want: []int{0, 0, 1, 2, 2, 3, 3, 4, 4, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := counting_sort(tt.args.x, tt.args.k)
			if !reflect.DeepEqual(r, tt.want) {
				t.Errorf(`
				name: %v,
				act. %v
				exp. %v
			`, tt.name, tt.args.x, tt.want)
			}
		})
	}
}
