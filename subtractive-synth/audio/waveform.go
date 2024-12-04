package audio

import (
	"math"
	"math/rand"
)

type WaveformType string

const (
	WaveformSine     WaveformType = "sine"
	WaveformSquare   WaveformType = "square"
	WaveformTriangle WaveformType = "triangle"
	WaveformSawtooth WaveformType = "sawtooth"
	WaveformPulse    WaveformType = "pulse"
	WaveformNoise    WaveformType = "noise"
)

func GenerateWaveform(waveform WaveformType, phase, pulseWidth float64) float64 {
	switch waveform {
	case WaveformSine:
		// Equation for a sine wave is y = sin(2 * pi * phase)
		return math.Sin(2 * math.Pi * phase)
	case WaveformSquare:
		// Equation for a square wave is y = sign(sin(2 * pi * phase)) but we can do it piecewise
		if phase < 0.5 {
			return 1.0
		}
		return -1.0
	case WaveformTriangle:
		// Equation for a triangle wave is y = 4 * abs(round(phase) - phase) - 1
		// Rather 2|2(x-floor(1/2+x))| - 1
		return 4*math.Abs(phase-0.5) - 1.0
	case WaveformSawtooth:
		// Equation for a sawtooth wave is y = 2 * phase - floor(phase+0.5)
		return 2*phase - math.Floor(phase+0.5)
	case WaveformPulse:
		// Equation for a pulse wave is y = sign(sin(2 * pi * phase))
		if phase < pulseWidth {
			return 1.0
		}
		return -1.0
	case WaveformNoise:
		return rand.Float64()*2 - 1.0
	default:
		return 0.0
	}
}
