package main

import "testing"

func Test_sum(t *testing.T) {
	type args struct {
		integers []int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"positive integers",
			args{integers: []int64{0, 1, 1, 2, 3, 5}},
			12,
		}, {
			"positive integers",
			args{integers: []int64{0, 1, 1, 2, 3}},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.integers...); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
