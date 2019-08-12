package main

import "testing"

func Test_fb(t *testing.T) {
	type args struct {
		vars []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				vars: []int{1, 3},
			},
			want: "Fizz",
		},
		{
			args: args{
				vars: []int{1, 3, 5},
			},
			want: "Fizz Buzz",
		},
		{
			args: args{
				vars: []int{1, 3, 5, 15},
			},
			want: "Fizz Buzz Fizz Buzz",
		},
		{
			args: args{
				vars: []int{1, 2, 3, 5, 6, 15},
			},
			want: "Fizz Buzz Fizz Fizz Buzz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fb(tt.args.vars); got != tt.want {
				t.Errorf("fb() = %v, want %v", got, tt.want)
			}
		})
	}
}
