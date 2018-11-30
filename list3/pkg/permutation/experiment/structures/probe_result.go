package structures

type ProbeResult struct {
	FixedPoints *ProbeDataset
	Cycles      *ProbeDataset
	Records     *ProbeDataset
}

func NewProbeResult(size int) *ProbeResult {
	probeResult := new(ProbeResult)
	probeResult.FixedPoints = NewProbeDataset(size)
	probeResult.Cycles = NewProbeDataset(size)
	probeResult.Records = NewProbeDataset(size)
	return probeResult
}
