package audio

import (
	"sync"

	"subtractive-synth/pkg/util"
)

// As everywhere else, a struct is made so a  the LFO (or oscillator, ADSR, filter, etc)
// can be created and used in the main program.
type LFO struct {
	Waveform   WaveformType
	Frequency  float64
	Amplitude  float64
	Phase      float64
	SampleRate float64
	Mutex      sync.Mutex
}

func NewLFO(waveform WaveformType, frequency, amplitude float64) *LFO {
	return &LFO{ // this function creates a new lfo by populating the struct above
		Waveform:   waveform,
		Frequency:  frequency,
		Amplitude:  amplitude,
		SampleRate: util.SampleRate, // 44100 Hz
		Phase:      0.0,
	}
}

// this explains itself, it sets the frequency for the lfo..

func (lfo *LFO) SetFrequency(freq float64) {
	lfo.Mutex.Lock()
	defer lfo.Mutex.Unlock()
	lfo.Frequency = freq
}

func (lfo *LFO) GenerateSample() float64 {
	lfo.Mutex.Lock()
	defer lfo.Mutex.Unlock()

	sample := GenerateWaveform(lfo.Waveform, lfo.Phase, 0.5) * lfo.Amplitude
	// that 0.5 is there from Generate Wave form. Needs clarification
	phaseIncrement := lfo.Frequency / lfo.SampleRate
	lfo.Phase += phaseIncrement
	if lfo.Phase >= 1.0 {
		lfo.Phase -= 1.0
	}
	return sample
}
