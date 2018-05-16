package tree

import "testing"

func TestTree(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ok",
			args: args{
				path: "./tmp",
			},
		},
	}
	for _, tt := range tests {
		Tree(tt.args.path)
	}
}
