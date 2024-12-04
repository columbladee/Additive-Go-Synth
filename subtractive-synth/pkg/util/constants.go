package util

import "math"

const (

	// Audio Settings
	SampleRate = 44100.0 // CD-quality
	BufferSize = 512     //Samples per buffer

	//Synth Settings

	MaxPolyphony   = 8
	NumOscillators = 2

	// Amplitude Settings
	MaxAmplitude = 1.0
	MinAmplitude = -1.0

	// Math Constant(s)
	TwoPi = 2*math.Pi

	// Midi Settings
	MidiNoteA4    = 60 // A4
	SemitoneRatio = 1.0594630943592952646 // from math.Pow(2, 1/12)

	//Envelope stage

	EnvelopeAttack = "attack"
	EnvelopeDecay  = "decay"

	EnvelopeSustain = "sustain"
	EnvelopeRelease = "release"

	// Waveform Types
	WaveformSine     = "sine"
	WaveformSquare   = "square"
	WaveformSawtooth = "sawtooth"
	WaveformTriangle = "triangle"
	WaveformNoise    = "noise"
)
