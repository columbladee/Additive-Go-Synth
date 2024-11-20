package util

import (
	"encoding/json"
	"os"
)
//again, more will be added to this config as more will be added to constants.go
type Config struct {
	SampleRate	int	`json:"sample_rate"`
	Polyphony	int	`json:"polyphony"`
	DutyCycle	float64	`json:"duty_cycle"`
}

// Lacks Setters and Getters, to be added after waveform and oscillator are in a better spot

func DefaultConfig() {
	return Config{
		SampleRate:	DefaultSampleRate, // from constants.go
		Polyphony:	MaxPolyphony,
		DutyCycle:	DefaultDutyCycle,	
	}
}


func LoadConfig(filePath string) (Config, error)  {
	file , err := os.Open(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist so return defaultconfig
			return DefaultConfig(), nil
		} else { // Some other error
			return Config{}, err
		}
	}
	// Stop it from closing
	defer file.Close()
	// set variable
	var config Config
	//Deal with JSON 
	// Somehow if there's a json decoding arrow commit suicide and apply for jobs in that order
	err = json.NewDecoder(file).Decode(&config)
	//ideally return config, nil
	// But guess what we need to handle those errors if they exist anyways, buddy
	if err != nil {
		//JSON Decoding likely
		return Config{}, err
	}
	// The ideal case
	return config, nil
}



