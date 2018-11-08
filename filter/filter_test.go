package filter

import (
	"reflect"
	"testing"
	"time"
)

func TestTimeSeries_Drop(t *testing.T) {
	type fields struct {
		interval    float64
		defaultData []TimeData
	}
	type args struct {
		base time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TimeSeries
	}{
		{
			fields: fields{
				interval: 3,
				defaultData: []TimeData{
					TimeData{duration: time.Duration(3 * time.Second)},
					TimeData{duration: time.Duration(4 * time.Second)},
					TimeData{duration: time.Duration(5 * time.Second)},
				},
			},
			args: args{
				base: time.Duration(6 * time.Second),
			},
			want: &TimeSeries{
				Data: []TimeData{
					TimeData{duration: time.Duration(3 * time.Second)},
					TimeData{duration: time.Duration(4 * time.Second)},
					TimeData{duration: time.Duration(5 * time.Second)},
				},
			},
		},
		{
			fields: fields{
				interval: 3,
				defaultData: []TimeData{
					TimeData{duration: time.Duration(3 * time.Second)},
					TimeData{duration: time.Duration(4 * time.Second)},
					TimeData{duration: time.Duration(5 * time.Second)},
				},
			},
			args: args{
				base: time.Duration(7 * time.Second),
			},
			want: &TimeSeries{
				Data: []TimeData{
					TimeData{duration: time.Duration(4 * time.Second)},
					TimeData{duration: time.Duration(5 * time.Second)},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := NewTimeSeries(tt.fields.interval)
			ts.Data = tt.fields.defaultData

			ts.Drop(tt.args.base)
			if !reflect.DeepEqual(ts.Data, tt.want.Data) {
				t.Fatalf("want %#v, but %#v", ts.Data, tt.want.Data)
			}
		})
	}
}
