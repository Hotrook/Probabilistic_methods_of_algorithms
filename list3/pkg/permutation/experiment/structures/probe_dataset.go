package structures

import "math"

type ProbeDataset struct {
	Points                  []int
	average                 float64
	concentration           float64
	standardDeviation       float64
	averageCalculated       bool
	concentrationCalculated bool
}

func NewProbeDataset(size int) *ProbeDataset {
	probeDataset := new(ProbeDataset)
	probeDataset.Points = make([]int, size)
	return probeDataset
}

func (pd *ProbeDataset) GetAverage() float64 {
	pd.calculateAverageIfNotCalculated()
	return pd.average
}

func (pd *ProbeDataset) calculateAverageIfNotCalculated() {
	if !pd.averageCalculated {
		sum := 0
		for _, element := range pd.Points {
			sum += element
		}
		pd.average = float64(sum) / float64(len(pd.Points))
		pd.averageCalculated = true
	}
}

func (pd *ProbeDataset) GetConcentration() float64 {
	pd.calculateConcentrationIfNotCalculated()
	return pd.concentration
}

func (pd *ProbeDataset) calculateConcentrationIfNotCalculated() {
	pd.calculateAverageIfNotCalculated()
	if !pd.concentrationCalculated {
		m := 0.0
		standardDeviation := 0.0
		for _, element := range pd.Points {
			m += math.Pow(float64(element)-pd.average, 4.0)
			standardDeviation += math.Pow(float64(element)-pd.average, 2.0)
		}
		size := float64(len(pd.Points))
		pd.concentration = (m / size) / math.Pow(standardDeviation/size, 2.0)
		pd.concentrationCalculated = true
	}
}
