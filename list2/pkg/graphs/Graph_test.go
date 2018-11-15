package graphs

import (
	"testing"
)

func TestGraph_addEdge(t *testing.T) {

	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		n    int
		args args
	}{
		{
			"Test 1",
			5,
			args{4, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := NewGraph(tt.n)
			graph.addEdge(tt.args.i, tt.args.j)
			haveEdge := graph.isEdge(tt.args.i, tt.args.j)

			if haveEdge != true {
				t.Error("Edge is not added.")
			}
		})
	}
}

func TestGraph_twoLargestComponents(t *testing.T) {
	type fields struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name       string
		fields     fields
		component1 int
		component2 int
	}{
		{
			"Test component 1",
			fields{5, [][]int{
				{0, 1, 0, 0, 0},
				{1, 0, 1, 0, 0},
				{0, 1, 0, 0, 0},
				{0, 0, 0, 0, 1},
				{0, 0, 0, 1, 0},
			}},
			3,
			2,
		},
		{
			"Test component 2",
			fields{5, [][]int{
				{0, 1, 0, 0, 0},
				{1, 0, 1, 0, 0},
				{0, 1, 0, 1, 0},
				{0, 0, 1, 0, 1},
				{0, 0, 0, 1, 0},
			}},
			5,
			0,
		},
		{
			"Test component 2",
			fields{5, [][]int{
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			}},
			1,
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := Graph{
				n:     tt.fields.n,
				edges: tt.fields.edges,
			}
			got, got1 := graph.TwoLargestComponents()
			if got != tt.component1 {
				t.Errorf("Graph.TwoLargestComponents() got = %v, component1 %v", got, tt.component1)
			}
			if got1 != tt.component2 {
				t.Errorf("Graph.TwoLargestComponents() got1 = %v, component1 %v", got1, tt.component2)
			}
		})
	}
}
