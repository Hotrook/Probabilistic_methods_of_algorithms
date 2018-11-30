package structures

type ExperimentDataset struct {
	Average       []float64
	Concentration []float64
}

func NewExperimentDataset(size int) *ExperimentDataset {
	experimentDataset := new(ExperimentDataset)
	experimentDataset.Average = make([]float64, size)
	experimentDataset.Concentration = make([]float64, size)
	return experimentDataset
}
