package sorting

import (
	"reflect"
	"testing"
)

func TestYaroslavskyDoublePivotQsort_sortArr(t *testing.T) {
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
