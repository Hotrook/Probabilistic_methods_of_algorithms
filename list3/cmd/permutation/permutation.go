package main

import (
	"flag"
	"fmt"
	"github.com/Hotrook/Probabilistic_methods_of_algorithms/list3/cmd/permutation/configuration"
	. "github.com/Hotrook/Probabilistic_methods_of_algorithms/list3/pkg/permutation"
	. "github.com/Hotrook/Probabilistic_methods_of_algorithms/list3/pkg/permutation/experiment/structures"
	"log"
	"math/rand"
	"os"
)

var n = flag.Int("n", configuration.GetInstance().GetDefaultN(), "This is max permutation size")
var step = flag.Int("step", configuration.GetInstance().GetDefaultStep(), "This is value by which size of permutation in experiment will change")
var start = flag.Int("start", configuration.GetInstance().GetDefaultStart(), "This is minimal value for size of permutation in experiment")
var probes = flag.Int("probes", configuration.GetInstance().GetDefaultProbes(), "This is minimal value for probes number for each size")

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	flag.Parse()
	n := *n
	step := *step
	start := *start
	probes := *probes
	filename := fmt.Sprintf(configuration.GetInstance().GetFilename(), start, step, n)

	log.Printf("Runing experiment with following data: [n=%d], [step=%d], [start=%d], [probes=%d]\n", n, step, start, probes)

	experimentResult := runExperiment(start, step, n, probes)
	saveToFile(experimentResult, filename, start, step, n)
}

func saveToFile(result *ExperimentResult, fileName string, start int, step int, n int) {

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Cannot create a file.")
	}
	defer file.Close()

	for size := start; size <= n; size += step {
		index := size/step - 1
		fmt.Fprintf(file, "%d", size)
		fmt.Fprintf(file, ";%f", result.FixedPoints.Average[index])
		fmt.Fprintf(file, ";%f", result.FixedPoints.Concentration[index])
		fmt.Fprintf(file, ";%f", result.Cycles.Average[index])
		fmt.Fprintf(file, ";%f", result.Cycles.Concentration[index])
		fmt.Fprintf(file, ";%f", result.Records.Average[index])
		fmt.Fprintf(file, ";%f", result.Records.Concentration[index])
		fmt.Fprintf(file, "\n")
	}
}

func runExperiment(start int, step int, n int, probes int) *ExperimentResult {
	datasetSize := n / step

	experimentResult := NewExperimentResult(datasetSize)

	for size := start; size <= n; size += step {
		log.Printf("Processing size %d", size)
		index := size/step - 1
		probeResult := runExperimentForSize(size, probes)
		experimentResult.Add(index, probeResult)
	}

	return experimentResult
}

func runExperimentForSize(size, probes int) *ProbeResult {
	var probeResult = NewProbeResult(probes)

	for probe := 0; probe < probes; probe++ {
		perm := rand.Perm(size)
		probeResult.Records.Points[probe] = CalculateRecords(perm)
		probeResult.Cycles.Points[probe] = CalculateCycles(perm)
		probeResult.FixedPoints.Points[probe] = CalculateFixedPoints(perm)
	}

	return probeResult
}
