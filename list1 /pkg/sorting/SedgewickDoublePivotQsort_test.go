package sorting

import (
	"reflect"
	"testing"
)

func TestSedgewickDoublePivotQsort_sort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qsort := &SedgewickDoublePivotQsort{}
			qsort.Sort(tt.args.arr)

			if !reflect.DeepEqual(tt.expected, tt.args.arr) {
				t.Error("Array is not sorted: %w", tt.args.arr)
			}
		})
	}
}
