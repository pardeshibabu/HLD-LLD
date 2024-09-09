package statedp

import "fmt"

// State Implementation
type TvState interface {
	State()
}

// Concrete implementation of State
type TvOn struct{}

func (to *TvOn) State() {
	fmt.Println("Tv is ON!")
}

type TvOff struct{}

func (to *TvOff) State() {
	fmt.Println("Tv is Off!")
}

// Context Implementation
type StateContext struct {
	CurrentTvState TvState
}

// Constructor Method to get the COntext Object
func GetContext() *StateContext {
	return &StateContext{
		CurrentTvState: &TvOff{},
	}
}

// Getter and Setter Method Implementation
func (sc *StateContext) SetState(ts TvState) {
	sc.CurrentTvState = ts
}

func (sc *StateContext) GetState() {
	sc.CurrentTvState.State()
}
