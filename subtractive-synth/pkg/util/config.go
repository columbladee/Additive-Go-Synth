package util

import (

	"encoding/json"
	"os"
	"log"
)

type Config struct {
	SampleRate	float64	`json:"sample_rate"`
	BufferSize	int	`json:"buffer_size"`
	NumOscillators	int	`json:"num_oscillators"`
	MidiInput	string	
}


// Config struct samplerate buffersize numosc midiinput

//loadconfig
