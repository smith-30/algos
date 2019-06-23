package lib

import "testing"

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
