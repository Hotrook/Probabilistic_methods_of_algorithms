package structures

type ExperimentResult struct {
	FixedPoints *ExperimentDataset
	Cycles      *ExperimentDataset
	Records     *ExperimentDataset
}

func NewExperimentResult(size int) *ExperimentResult {
	experimentResult := new(ExperimentResult)
	experimentResult.Records = NewExperimentDataset(size)
	experimentResult.Cycles = NewExperimentDataset(size)
	experimentResult.FixedPoints = NewExperimentDataset(size)
	return experimentResult
}

func (result ExperimentResult) Add(index int, probeResult *ProbeResult) {
	result.FixedPoints.Average[index] = probeResult.FixedPoints.GetAverage()
	result.FixedPoints.Concentration[index] = probeResult.FixedPoints.GetConcentration()
	result.Cycles.Average[index] = probeResult.Cycles.GetAverage()
	result.Cycles.Concentration[index] = probeResult.Cycles.GetConcentration()
	result.Records.Average[index] = probeResult.Records.GetAverage()
	result.Records.Concentration[index] = probeResult.Records.GetConcentration()
}
