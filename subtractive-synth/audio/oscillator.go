package audio

import (
	"subtractive-synth/pkg/util"
	"sync"
)

type Oscillator struct {
	Waveform   WaveformType
	Frequency  float64
	Amplitude  float64
	PulseWidth float64
	Phase      float64
	SampleRate float64
	Mutex      sync.Mutex
}

func NewOscillator(waveform WaveformType, frequency float64, amplitude float64, pulseWidth float64) *Oscillator {
	return &Oscillator{
		Waveform:   waveform,
		Frequency:  frequency,
		Amplitude:  amplitude,
		PulseWidth: pulseWidth,
		Phase:      0.0,
		SampleRate: util.SampleRate,
	}
}

func (osc *Oscillator) SetFrequency(freq float64) {
	osc.Mutex.Lock()
	defer osc.Mutex.Unlock()
	osc.Frequency = freq
}

func (osc *Oscillator) SetAmplitude(amp float64) {
	osc.Mutex.Lock()
	defer osc.Mutex.Unlock()
	osc.Amplitude = amp
}

func (osc *Oscillator) SetPulseWidth(pw float64) {
	osc.Mutex.Lock()
	defer osc.Mutex.Unlock()
	osc.PulseWidth = pw
}

func (osc *Oscillator) GenerateSample() float64 {
	osc.Mutex.Lock()
	defer osc.Mutex.Unlock()

	sample := GenerateWaveform(osc.Waveform, osc.Phase, osc.PulseWidth) * osc.Amplitude

	phaseIncrement := osc.Frequency / osc.SampleRate

	osc.Phase += phaseIncrement

	if osc.Phase >= 1.0 {
		osc.Phase -= 1.0
	}
	return sample
}
