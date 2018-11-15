package main

import "math"

func main() {
	n := 1000
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
		//component1sum := 0
		//component2sum := 0

		for probe := 0; probe < probes; probe++ {
			//graph := generateRandomGraph(n)
		}
	}

}
