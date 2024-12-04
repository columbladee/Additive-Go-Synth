package audio

import (
	"subtractive-synth/pkg/util"
)

type FilterType int

const (
	LowPass  FilterType = iota // 1
	HighPass                   // 2
	BandPass                   // 3
	Notch                      // 4
)

type Filter struct {
	FilterType FilterType
	Cutoff     float64
	Resonance  float64 // Q Factor - controls emphasis of frequences around the cutoff frequency. makes a "resonant" or "ringing" sound

	sampleRate float64

	buf0, buf1 float64
	// Buffers for the filter -- used to store previous values
	// The point of these buffers is to store the previous values of the filter so that we can use them in the next iteration
}

func NewFilter(filterType FilterType, cutoff, resonance float64) *Filter {
	return &Filter{
		FilterType: filterType,
		Cutoff:     cutoff,
		Resonance:  resonance,
		sampleRate: util.SampleRate, // 44100 Hz
	}
}

func (f *Filter) ProcessSample(input float64) float64 {
	f.buf0 += f.Cutoff * (input - f.buf0 + f.Resonance*(f.buf0-f.buf1)) // buf0 is updated. adds a value based on the cutoff,difference
	// between the input signal and current value of f.buf0, and the resonance value times the difference between f.buf0 and f.buf1
	//The resonance effect added is based on the difference between the two buffers
	f.buf1 += f.Cutoff * (f.buf0 - f.buf1)

	switch f.FilterType {
	case LowPass:
		return f.buf1
	case HighPass:
		return input - f.buf0
	case BandPass:
		return f.buf0 - f.buf1
	case Notch:
		return input - (f.buf0 - f.buf1)
	default:
		return input
	}
}

// Writeup / explanation of the filter is necessary, this isn't clear to someone who reads it
