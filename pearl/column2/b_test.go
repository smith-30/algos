package column2

import (
	"reflect"
	"testing"
)

func TestRotateTaskB(t *testing.T) {
	type args struct {
		src   []string
		count int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "ok_0",
			args: args{
				src:   []string{"a", "b", "c", "d", "e"},
				count: 0,
			},
			want: []string{"a", "b", "c", "d", "e"},
		},
		{
			name: "ok_1",
			args: args{
				src:   []string{"a", "b", "c", "d", "e"},
				count: 1,
			},
			want: []string{"b", "c", "d", "e", "a"},
		},
		{
			name: "ok_2",
			args: args{
				src:   []string{"a", "b", "c", "d", "e"},
				count: 2,
			},
			want: []string{"c", "d", "e", "a", "b"},
		},
		{
			name: "ok_5",
			args: args{
				src:   []string{"a", "b", "c", "d", "e"},
				count: 5,
			},
			want: []string{"a", "b", "c", "d", "e"},
		},
		{
			name: "ok_6",
			args: args{
				src:   []string{"a", "b", "c", "d", "e"},
				count: 6,
			},
			want: []string{"b", "c", "d", "e", "a"},
		},
	}
	for _, tt := range tests {
		RotateTaskB(tt.args.src, tt.args.count)
		if !reflect.DeepEqual(tt.args.src, tt.want) {
			t.Errorf(`
			name: %v,
			act. %v
			exp. %v
		`, tt.name, tt.args.src, tt.want)
		}
	}
}

func TestNewRotateTaskB(t *testing.T) {
	type args struct {
		src   []string
		count int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "ok_0",
			args: args{
				src:   []string{"a", "b", "c", "d", "e"},
				count: 0,
			},
			want: []string{"a", "b", "c", "d", "e"},
		},
		{
			name: "ok_1",
			args: args{
				src:   []string{"a", "b", "c", "d", "e"},
				count: 1,
			},
			want: []string{"b", "c", "d", "e", "a"},
		},
		{
			name: "ok_2",
			args: args{
				src:   []string{"a", "b", "c", "d", "e"},
				count: 2,
			},
			want: []string{"c", "d", "e", "a", "b"},
		},
		{
			name: "ok_5",
			args: args{
				src:   []string{"a", "b", "c", "d", "e"},
				count: 5,
			},
			want: []string{"a", "b", "c", "d", "e"},
		},
		{
			name: "ok_6",
			args: args{
				src:   []string{"a", "b", "c", "d", "e"},
				count: 6,
			},
			want: []string{"b", "c", "d", "e", "a"},
		},
	}
	for _, tt := range tests {
		NewRotateTaskB(tt.args.src, tt.args.count)
		if !reflect.DeepEqual(tt.args.src, tt.want) {
			t.Errorf(`
			name: %v,
			act. %v
			exp. %v
		`, tt.name, tt.args.src, tt.want)
		}
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		src []string
		s   int
		e   int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "ok_1",
			args: args{
				src: []string{"a", "b", "c", "d"},
				s:   0,
				e:   1,
			},
			want: []string{"b", "a", "c", "d"},
		},
		{
			name: "ok_2",
			args: args{
				src: []string{"a", "b", "c", "d"},
				s:   0,
				e:   2,
			},
			want: []string{"c", "b", "a", "d"},
		},
		{
			name: "ok_3",
			args: args{
				src: []string{"a", "b", "c", "d"},
				s:   1,
				e:   2,
			},
			want: []string{"a", "c", "b", "d"},
		},
		{
			name: "ok_4",
			args: args{
				src: []string{"a", "b", "c", "d"},
				s:   1,
				e:   3,
			},
			want: []string{"a", "d", "c", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverse(tt.args.src, tt.args.s, tt.args.e)
			if !reflect.DeepEqual(tt.args.src, tt.want) {
				t.Errorf("\n%v, act %v, want %v", tt.name, tt.args.src, tt.want)
			}
		})
	}
}