package main

import (
	"fmt"
	"github.com/Hotrook/Probabilistic_methods_of_algorithms/list1/pkg/sorting"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"math/rand"
)

const n = 500
const algorithmsNumber = 3

var algoNames = []string{"Regular", "Sedgewick", "Yaroslavsky"}

func main() {

	probes := 100

	swaps := [algorithmsNumber][n]int{}
	comparisons := [algorithmsNumber][n]int{}
	algorithms := []sorting.Quicksort{&sorting.RegularQsort{}, &sorting.SedgewickDoublePivotQsort{}, &sorting.YaroslavskyDoublePivotQsort{}}

	makeExperiment(probes, algorithms, swaps, comparisons)

	generatePlot(comparisons, "comparisons.png", "comparisons")
	generatePlot(swaps, "swaps.png", "swaps")
}

func makeExperiment(probes int, algorithms []sorting.Quicksort, swaps [algorithmsNumber][n]int, comparisons [algorithmsNumber][n]int) {
	for size := 0; size < n; size++ {
		fmt.Println("%5w / %5w \r", size, n)
		makeProbesForSize(probes, size, algorithms, swaps, comparisons)
		for algoId := 0; algoId <= 2; algoId++ {
			swaps[algoId][size] /= n
			comparisons[algoId][size] /= n
		}
	}
}

func makeProbesForSize(probes int, size int, algorithms []sorting.Quicksort, swaps [algorithmsNumber][n]int, comparisons [algorithmsNumber][n]int) {
	for probe := 1; probe <= probes; probe++ {
		array := generateRandomPermutation(size + 1)
		toPass := make([]int, size+1)
		for id, algorithm := range algorithms {
			copy(toPass, array)
			s, c := algorithm.Sort(toPass)
			swaps[id][size] += s
			comparisons[id][size] += c
		}
	}
}

func generatePlot(swaps [algorithmsNumber][n]int, filename, title string) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = title

	err = plotutil.AddLinePoints(p,
		algoNames[0], createDataSet(swaps, 0),
		algoNames[1], createDataSet(swaps, 1),
		algoNames[2], createDataSet(swaps, 2))

	if err != nil {
		panic(err)
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, filename); err != nil {
		panic(err)
	}
}

func createDataSet(swaps [algorithmsNumber][n]int, algoId int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		pts[i].X = float64(i)
		pts[i].Y = float64(swaps[algoId][i])
	}
	return pts
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
