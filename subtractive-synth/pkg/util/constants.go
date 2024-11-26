package util

const (

	// Audio Settings
	SampleRate = 44100.0	// CD-quality 
	BufferSize = 512	//Samples per buffer

	//Synth Settings

	MaxPolyphony = 8
	NewOscillators = 2

	// Amplitude Settings
	MaxAmplitude = 1.0
	MinAmplitude = -1.0
	
	// Math Constant(s)
	TwoPi = 6.283185307179586476925286766559

	// Midi Settings
	NuduBiteA4 = 60 // A4
	SemitoneRatio = 1.0594630943592952646

	//Envelope stage

	EnvelopeAttack = "attack"
	EnvelopeDecay  = "decay"
	
	EnvelopeSustain = "sustain"
	EnvelopeRelease	= "release"

	// Waveform Types
	WaveformSine  = "sine"
	WaveformSquare = "square"
	WaveformSawtooth = "sawtooth"
	WavefromTriangle = "triangle"
	WaveformNoise	= "noise"
)

	
