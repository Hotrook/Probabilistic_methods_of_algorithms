package permutation

import (
	"testing"
)

func Test_CalculateFixedPoints(t *testing.T) {
	type args struct {
		permutation []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"All fixed points.",
			args{[]int{0, 1, 2, 3, 4}},
			5,
		},
		{
			"One fixed point.",
			args{[]int{4, 3, 2, 1, 0}},
			1,
		},
		{
			"Three fixed point.",
			args{[]int{0, 3, 2, 1, 4}},
			3,
		},
		{
			"No fixed point.",
			args{[]int{1, 2, 3, 4, 0}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateFixedPoints(tt.args.permutation); got != tt.want {
				t.Errorf("CalculateFixedPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CalculateCycles(t *testing.T) {
	type args struct {
		permutation []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"5 cycles",
			args{[]int{0, 1, 2, 3, 4}},
			5,
		},
		{
			"3 cycles",
			args{[]int{4, 3, 2, 1, 0}},
			3,
		},
		{
			"4 cycles",
			args{[]int{0, 3, 2, 1, 4}},
			4,
		},
		{
			"1 cycle",
			args{[]int{1, 2, 3, 4, 0}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateCycles(tt.args.permutation); got != tt.want {
				t.Errorf("CalculateCycles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CalculateRecords(t *testing.T) {
	type args struct {
		permutation []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"5 records",
			args{[]int{0, 1, 2, 3, 4}},
			5,
		},
		{
			"1 record",
			args{[]int{4, 3, 2, 1, 0}},
			1,
		},
		{
			"3 records",
			args{[]int{0, 3, 2, 1, 4}},
			3,
		},
		{
			"4 records",
			args{[]int{1, 2, 3, 4, 0}},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateRecords(tt.args.permutation); got != tt.want {
				t.Errorf("CalculateRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}
