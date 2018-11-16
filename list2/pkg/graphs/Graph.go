package graphs

import (
	"math/rand"
	"sort"
	"time"
)

type Graph struct {
	n     int
	edges [][]int
}

func NewGraph(n int) *Graph {
	graph := new(Graph)
	graph.n = n
	graph.edges = make([][]int, n)
	for i := range graph.edges {
		graph.edges[i] = make([]int, n)
	}
	return graph
}

func (graph Graph) addEdge(i, j int) {
	graph.edges[i][j] = 1
	graph.edges[j][i] = 1
}

func (graph Graph) isEdge(i int, j int) bool {
	return graph.edges[i][j] == 1
}

func (graph Graph) TwoLargestComponents() (int, int) {

	nodeComponent, componentNumber := graph.divideNodesIntoComponents()
	components := make([]int, componentNumber)
	for node := range nodeComponent {
		components[nodeComponent[node]-1]++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(components)))

	if componentNumber > 1 {
		return components[0], components[1]
	} else {
		return components[0], 0
	}
}

func (graph Graph) divideNodesIntoComponents() ([]int, int) {
	nodeComponent := make([]int, graph.n)
	componentNumber := 0

	for node := 0; node < graph.n; node++ {
		if nodeComponent[node] == 0 {
			componentNumber++
			graph.bfs(node, nodeComponent, componentNumber)
		}
	}

	return nodeComponent, componentNumber
}

func (graph Graph) bfs(node int, nodeComponent []int, component int) {
	queue := make([]int, 0)
	queue = append(queue, node)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		nodeComponent[current] = component

		for neighbour := range graph.edges[current] {
			if graph.edges[current][neighbour] == 1 && nodeComponent[neighbour] == 0 {
				queue = append(queue, neighbour)
			}
		}
	}
}

func GenerateRandomGraph(n int, p float64) *Graph {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	graph := NewGraph(n)
	for i := range graph.edges {
		for j := i + 1; j < n; j++ {
			if r1.Float64() <= p {
				graph.addEdge(i, j)
			}
		}
	}

	return graph
}
