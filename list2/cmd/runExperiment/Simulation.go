package main

import (
	"flag"
	"fmt"
	"github.com/Hotrook/Probabilistic_methods_of_algorithms/list2/cmd/runExperiment/configuration"
	"github.com/Hotrook/Probabilistic_methods_of_algorithms/list2/pkg/graphs"
	"log"
	"math"
	"os"
)

type probabilityFunction func(int) float64

var n = flag.Int("n", 1000, "This is max size for graph in experiment")
var step = flag.Int("step", 50, "This is value by which size for graph in experiment will change")
var start = flag.Int("start", 50, "This is minimal value for size for graph in experiment")

func main() {

	flag.Parse()
	n := *n
	step := *step
	start := *start

	probes := 100

	functions := []probabilityFunction{first, second, third, fourth, fifth, sixth}

	first, second := generateDataForCharts(functions, n, step, start, probes)

	saveDatasetToFile(first, fmt.Sprintf(configuration.GetInstance().GetFstComponentFileName(), start, step, n), step)
	saveDatasetToFile(second, fmt.Sprintf("secondComponent_%d_%d_%d.txt", start, step, n), step)
}

func first(n int) float64 {
	return 1.0 / (2.0 * float64(n))
}

func second(n int) float64 {
	return (1.0 / float64(n)) - (math.Pow(float64(n), 0.1) / math.Pow(float64(n), 4.0/3.0))
}

func third(n int) float64 {
	return (1.0 / float64(n)) - (2.0 / math.Pow(float64(n), 4.0/3.0))
}

func fourth(n int) float64 {
	return (1.0 / float64(n)) + (2.0 / math.Pow(float64(n), 4.0/3.0))
}

func fifth(n int) float64 {
	return (1.0 / float64(n)) + (math.Pow(float64(n), 0.1) / math.Pow(float64(n), 4.0/3.0))
}

func sixth(n int) float64 {
	return 2.0 / float64(n)
}

func saveDatasetToFile(dataSet [][]float64, fileName string, step int) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Cannot create a file.")
	}
	defer file.Close()

	for i := range dataSet {
		size := (i + 1) * step
		fmt.Fprintf(file, "%d", size)
		for j := range dataSet[i] {
			fmt.Fprintf(file, ";%.2f", dataSet[i][j])
		}
		fmt.Fprintf(file, "\n")
	}
}

func generateDataForCharts(probabilities []probabilityFunction, n, step, start, probes int) ([][]float64, [][]float64) {
	firstCompResult := make([][]float64, n/step)
	secondCompResult := make([][]float64, n/step)

	for i := range firstCompResult {
		firstCompResult[i] = make([]float64, len(probabilities))
		secondCompResult[i] = make([]float64, len(probabilities))
	}

	for size := start; size <= n; size += step {
		i := size/step - 1
		fmt.Printf("\rWorking on size %d", size)
		for p := range probabilities {
			componentSize1, componentSize2 := makeExperiment(probes, size, probabilities[p](size))
			firstCompResult[i][p] = componentSize1
			secondCompResult[i][p] = componentSize2
		}
	}

	return firstCompResult, secondCompResult
}

func makeExperiment(probes int, n int, probability float64) (float64, float64) {
	component1sum := 0
	component2sum := 0

	for probe := 0; probe < probes; probe++ {
		graph := graphs.GenerateRandomGraph(n, probability)
		c1, c2 := graph.TwoLargestComponents()
		component1sum += c1
		component2sum += c2
	}

	componentAverage1 := float64(component1sum) / float64(probes)
	componentAverage2 := float64(component2sum) / float64(probes)

	return componentAverage1, componentAverage2
}
