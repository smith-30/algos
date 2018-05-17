package column2

import (
	"reflect"
	"testing"
)

func TestGetAnagram(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok_1",
			args: args{
				key: "post",
			},
			wantErr: false,
		},
		{
			name: "ok_2",
			args: args{
				key: "ostp",
			},
			wantErr: false,
		},
		{
			name: "ok_3",
			args: args{
				key: "stpo",
			},
			wantErr: false,
		},
		{
			name: "ok_4",
			args: args{
				key: "stpopp",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := sortStr(tt.args.key)
			d := getDict()
			if want, ok := d[k]; !ok {
				if !tt.wantErr {
					panic(notFoundErr)
				}
			} else {
				got, err := GetAnagram(tt.args.key)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetAnagram() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("GetAnagram() = %v, want %v", got, want)
				}
			}
		})
	}
}
