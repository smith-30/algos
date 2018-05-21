package column3

import "testing"

func Test_showSegment(t *testing.T) {
	type args struct {
		src [5]byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ok_0",
			args: args{
				src: [5]byte{255, 2, 3, 8, 255},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			showSegment(tt.args.src)
		})
	}
}
