package audio

import "subtractive-synth/pkg/util"

type ADSRState int // current state of the ADSR envelope

const (
	Attack  ADSRState = iota // 1
	Decay                    // 2
	Sustain                  //3
	Release                  //4
	Idle                     //5
)

type ADSR struct {
	AttackTime   float64
	DecayTime    float64
	SustainLevel float64
	ReleaseTime  float64

	State       ADSRState
	CurrentTime float64
	Output      float64
}

func NewADSR(attack, decay, sustain, release float64) *ADSR {
	return &ADSR{
		AttackTime:   attack,
		DecayTime:    decay,
		SustainLevel: sustain,
		ReleaseTime:  release,
		State:        Idle,
		Output:       0.0,
	}
}

func (env *ADSR) NoteOn() {
	env.State = Attack
	env.CurrentTime = 0.0
}

func (env *ADSR) NoteOff() {
	env.State = Release
	env.CurrentTime = 0.0
}

func (env *ADSR) GenerateSample() float64 {
	deltaTime := 1.0 / util.SampleRate

	switch env.State {
	case Attack:
		env.Output += deltaTime / env.AttackTime
		if env.Output >= 1.0 {
			env.Output = 1.0
			env.State = Decay
		}
	case Decay:
		env.Output -= deltaTime * (1.0 - env.SustainLevel) / env.DecayTime
		if env.Output <= env.SustainLevel || env.DecayTime == 0 {
			env.Output = env.SustainLevel
			env.State = Sustain
		}
	case Sustain:
		env.Output = env.SustainLevel
	case Release:
		env.Output -= deltaTime * env.SustainLevel / env.ReleaseTime
		if env.Output <= 0.0 || env.ReleaseTime == 0 {
			env.Output = 0.0
			env.State = Idle
		}
	case Idle:
		env.Output = 0.0
	}
	return env.Output
}
