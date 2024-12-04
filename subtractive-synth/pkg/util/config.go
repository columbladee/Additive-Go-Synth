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
	MidiInput	string	`json:"midi_input"`
}



//loadconfig

func LoadConfig(filename string) (*Config, error) {
	// content read with os.Readfile
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Failed to read config file: %v", err)
		return nil, err
	}

	var config Config
	// unmarshal json content
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Printf("Failed to parse config file: %v", err)
		return nil, err
	}

	return &config, nil

}

func DefaultConfig() *Config {
	return &Config {
		SampleRate:	SampleRate, // ALl of these are constants from constants.go
		BufferSize:	BufferSize,
		NumOscillators: NumOscillators,
		MidiInput:	"default",
	}
}


// add setters and getters later
