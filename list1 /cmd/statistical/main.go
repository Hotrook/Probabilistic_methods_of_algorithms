package main

import (
	"fmt"
	"github.com/Hotrook/Probabilistic_methods_of_algorithms/list1/pkg/sorting"
	"math"
	"math/rand"
)

type function func(n int) float64

const algorithmsNumber = 3
const n = 500
const probes = 1000

var algoNames = []string{"Regular", "Sedgewick", "Yaroslavsky"}

// Theoretical values for expected values of comparisons and swaps
var expcetedValues = [][]function{
	{
		func(n int) float64 {
			return 2.0*float64(n)*math.Log(float64(n)) - 1.51*float64(n)
		},
		func(n int) float64 {
			return 0.33*float64(n)*math.Log(float64(n)) - 0.25*float64(n)
		},
	},
	{
		func(n int) float64 {
			return 2.11*float64(n)*math.Log(float64(n)) - 2.57*float64(n)
		},
		func(n int) float64 {
			return 0.8*float64(n)*math.Log(float64(n)) - 0.3*float64(n)
		},
	},
	{
		func(n int) float64 {
			return 1.9*float64(n)*math.Log(float64(n)) - 2.46*float64(n)
		},
		func(n int) float64 {
			return 0.6*float64(n)*math.Log(float64(n)) - 0.08*float64(n)
		},
	},
}

func main() {

	swaps := [algorithmsNumber][probes]int{}
	comparisons := [algorithmsNumber][probes]int{}
	algorithm := []sorting.Quicksort{&sorting.RegularQsort{}, &sorting.SedgewickDoublePivotQsort{}, &sorting.YaroslavskyDoublePivotQsort{}}

	makeExperiment(algorithm, swaps, comparisons)

	generateStats(swaps, comparisons)

}

func makeExperiment(algorithm []sorting.Quicksort, swaps [algorithmsNumber][probes]int, comparisons [algorithmsNumber][probes]int) {
	for probe := 0; probe < probes; probe++ {
		array := generateRandomPermutation(n)
		toPass := make([]int, n)
		for algoId := 0; algoId <= 2; algoId++ {
			copy(toPass, array)
			s, c := algorithm[algoId].Sort(toPass)
			swaps[algoId][probe] = s
			comparisons[algoId][probe] = c
		}
	}
}

func generateStats(swaps [algorithmsNumber][probes]int, comparisons [algorithmsNumber][probes]int) {
	for alg := 0; alg < algorithmsNumber; alg++ {
		printStatsForAlgorithm(alg, comparisons, swaps)
	}
}

func printStatsForAlgorithm(alg int, comparisons, swaps [algorithmsNumber][probes]int) {
	fmt.Printf("Algorithm: %v\n", algoNames[alg])
	printStatsForSequence(alg, "\tC_n:", comparisons)
	printStatsForSequence(alg, "\tS_n:", swaps)
	printTheoreticalValues(alg)
	fmt.Println()
}

func printTheoreticalValues(alg int) {
	fmt.Printf("\tTheoretically: \n")
	fmt.Printf("\tC_n:\n")
	fmt.Printf("\t\tEX:  %v\n", expcetedValues[alg][0](n))
	fmt.Printf("\tS_n:\n")
	fmt.Printf("\t\tEX:  %v\n", expcetedValues[alg][1](n))
}

func printStatsForSequence(alg int, statName string, comparisons [algorithmsNumber][probes]int) {
	fmt.Println(statName)
	fmt.Printf("\t\tEX:  %v\n", calculateEX(comparisons, alg))
	fmt.Printf("\t\tVar: %v\n", calculateVar(comparisons, alg))
}

func calculateVar(seq [algorithmsNumber][probes]int, alg int) float64 {
	sum := 0.0
	ex := calculateEX(seq, alg)
	for i := 0; i < probes; i++ {
		sum += math.Pow(float64(seq[alg][i])-ex, 2)
	}
	return sum / probes
}

func calculateEX(seq [algorithmsNumber][probes]int, alg int) float64 {
	sum := 0.0
	for i := 0; i < probes; i++ {
		sum += float64(seq[alg][i])
	}
	return sum / probes
}

func generateRandomPermutation(n int) []int {
	result := make([]int, n)
	for index := range result {
		result[index] = index
	}
	rand.Shuffle(int(n), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return result
}
