package task1

// State state for task1
type State struct {
	No         int
	Position   string
	LeftClean  bool
	RightClean bool
	NextState  *State
}

var eighthState = &State{8, "right", true, true, nil}
var seventhState = &State{7, "left", true, true, nil}
var sixthState = &State{6, "right", true, false, eighthState}
var fifthState = &State{5, "left", true, false, sixthState}
var thirdState = &State{3, "left", false, true, seventhState}
var fourthState = &State{4, "right", false, true, thirdState}
var secondState = &State{2, "right", false, false, fourthState}
var firstState = &State{1, "left", false, false, fifthState}
var states = [...]*State{firstState, secondState, thirdState, fourthState, fifthState, sixthState, seventhState, eighthState}
