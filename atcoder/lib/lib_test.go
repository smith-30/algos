package lib

import (
	"reflect"
	"testing"
)

func TestLCM(t *testing.T) {
	type args struct {
		temp1 int
		temp2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				temp1: 4,
				temp2: 6,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LCM(tt.args.temp1, tt.args.temp2); got != tt.want {
				t.Errorf("LCM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nCr(t *testing.T) {
	type args struct {
		n int64
		r int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				n: 4,
				r: 2,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nCr(tt.args.n, tt.args.r); got != tt.want {
				t.Errorf("nCr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivisors(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				n: 10,
			},
			want: []int{1, 10, 2, 5},
		},
		{
			args: args{
				n: 1,
			},
			want: []int{1},
		},
		{
			args: args{
				n: 3,
			},
			want: []int{1, 3},
		},
		{
			args: args{
				n: 4,
			},
			want: []int{1, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Divisors(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Divisors() = %v, want %v", got, tt.want)
			}
		})
	}
}
