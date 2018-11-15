package main

import (
	"fmt"
	"github.com/Hotrook/Probabilistic_methods_of_algorithms/list2/pkg/graphs"
	"math"
)

func main() {
	n := 500
	nf := float64(n)
	probes := 100

	probabilities := []float64{
		1.0 / 2.0 * nf,
		(1.0 / nf) - (math.Pow(nf, 0.1) / math.Pow(nf, 4.0/3.0)),
		(1.0 / nf) - (2.0 / math.Pow(nf, 4.0/3.0)),
		(1.0 / nf) + (2.0 / math.Pow(nf, 4.0/3.0)),
		(1.0 / nf) + (math.Pow(nf, 0.1) / math.Pow(nf, 4.0/3.0)),
		2.0 / nf,
	}

	for p := range probabilities {
		component1sum := 0
		component2sum := 0

		for probe := 0; probe < probes; probe++ {
			graph := graphs.GenerateRandomGraph(n, probabilities[p])
			c1, c2 := graph.TwoLargestComponents()
			component1sum += c1
			component2sum += c2
		}

		componentAverage1 := float64(component1sum) / float64(probes)
		componentAverage2 := float64(component2sum) / float64(probes)
		fmt.Printf("For probablity %2d average size of  first component: %10f\n", p, componentAverage1)
		fmt.Printf("For probablity %2d average size of second component: %10f\n\n", p, componentAverage2)
	}

}
