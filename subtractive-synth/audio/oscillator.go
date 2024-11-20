package audio

type Oscillator struct {
	Frequency	float64
	Amplitude	float64
	Waveform	string
	DutyCycle	float64 //PWN
}

func NewOscillator(waveform string, frequency, amplitude dutyCycle float64) *Oscillator {
	return &Oscillator{
		Frequency:	frequency,
		Amplitude:	amplitude,
		Waveform:	waveform,
		DutyCycle:	dutyCycle,
	}
}


