package filter

import (
	"time"
)

type TimeData struct {
	duration time.Duration
	value    float64
}

type TimeSeries struct {
	Data     []TimeData
	required func(time.Duration, time.Duration) bool
}

func NewTimeSeries(i float64) *TimeSeries {
	return &TimeSeries{
		required: func(new, old time.Duration) bool {
			return (new.Seconds()-old.Seconds() <= i)
		},
	}
}

func (ts *TimeSeries) Drop(base time.Duration) {
	b := ts.Data[:0]
	for _, d := range ts.Data {
		if ts.required(base, d.duration) {
			b = append(b, d)
		}
	}
	ts.Data = b
}
