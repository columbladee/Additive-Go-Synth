package audio

type Oscillator struct {
	Frequency	float64
	Amplitude	float64
	Waveform	string
	DutyCycle	float64 //PWN
}

func NewOscillator(waveform St
