


Phase 1:

1.
Setup File Hirearchy:

Create directory structure
go mod init additive-synth

2. Core utilities
	pkg/util/constants.go    --- default constants
	pkg/util/config.go       ---- load and save config files (use default settings initially)
3. Waveform and oscillator:
	pkg/audio/wavetable.go ---- Generate waveforms (sine, square, triangle, sawtooth), normalize function
	pkg/audio/oscillator.go ---  Manages oscillator states and functionality
					handles: frequency, phase and amplitude of osc
					interactions with other components like LFO for modulation
					interfacing with waveform.go
					handles pulse width 


Phase 2: Audio Engine

4. Synth Engine
	a. audio/synth.go
		struct to manage voices
		function to mix osc outputs into a single audio stream
	b. audio/envelope.go
		create adsr envelope, apply to each voice
	c. audio/mixer.go
		manage max 8 voices dynamically, removing the oldest note when limits reached

5. Real-Time Audio Playback

PortAudio 
Test with hardcoded notes and simple waveforms


Phase 3 MIDI integration:

midi/midi.go using portMIDI
	-handle midi note on/off events to trigger oscillators
	-map velocity to amplitude and mod wheel to lfo params
	-test midi input by logging events and triggering sample tones


Phase 4: LFO / Modulation

7. Add LFO Support

	a. audio/lfo.go 
	b. integrate lfo.go into synth.go to modulate
		i. PWN duty cycle
		ii. Pitch or amplitude of oscillators


Phase 5: GUI Development


Phase 6: Testing and Optimziation

Phaze 7 Finalization

Deployment


use fyne package for cross platform binaries
	



