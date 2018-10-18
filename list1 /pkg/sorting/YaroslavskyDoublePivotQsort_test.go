package sorting

import (
	"reflect"
	"testing"
)

func TestYaroslavskyDoublePivotQsort_sortArr(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			"Regular RegularQsort",
			args{[]int{3, 2, 1}},
			[]int{1, 2, 3},
		},
		{
			"Already sorted",
			args{[]int{1, 2, 3}},
			[]int{1, 2, 3},
		},
		{
			"Regular RegularQsort",
			args{[]int{3, 4, 7, 5, 6, 9, 8, 2, 1}},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qsort := &YaroslavskyDoublePivotQsort{}
			qsort.Sort(tt.args.arr)

			if !reflect.DeepEqual(tt.expected, tt.args.arr) {
				t.Error("Array is not sorted: %w", tt.args.arr)
			}
		})
	}
}