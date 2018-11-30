package structures

import (
	"math"
	"testing"
)

func TestProbeDataset_GetAverage(t *testing.T) {
	type fields struct {
		size   int
		Points []int
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		epsilon float64
	}{
		{
			"",
			fields{5, []int{1, 2, 3, 4, 5}},
			3.0,
			0.0000001,
		},
		{
			"",
			fields{6, []int{1, 2, 3, 4, 5, 6}},
			3.5,
			0.0000001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pd := NewProbeDataset(tt.fields.size)
			pd.Points = tt.fields.Points
			if got := pd.GetAverage(); math.Abs(got-tt.want) > tt.epsilon {
				t.Errorf("ProbeDataset.GetAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProbeDataset_GetConcentration(t *testing.T) {
	type fields struct {
		size   int
		Points []int
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		epsilon float64
	}{
		{
			"",
			fields{5, []int{1, 2, 3, 4, 5}},
			1.7,
			0.0000001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pd := NewProbeDataset(tt.fields.size)
			pd.Points = tt.fields.Points
			if got := pd.GetConcentration(); got != tt.want {
				t.Errorf("ProbeDataset.GetConcentration() = %v, want %v", got, tt.want)
			}
		})
	}
}
