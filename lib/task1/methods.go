package task1

// GetNextMove return next move must be done
func (state State) GetNextMove() string {
	if state.NextState == nil {
		return ""
	}

	move := ""
	if state.Position == "left" && !state.LeftClean {
		move += " suck "
	} else if state.Position == "right" && !state.RightClean {
		move += " suck "
	}
	if state.Position != state.NextState.Position {
		move += " " + state.NextState.Position + " "
	}

	return move
}

// GetStep get steps
func GetStep(position string, leftClean bool, rightClean bool) *State {
	var currentState *State
	for _, value := range states {
		if value.Position == position && value.LeftClean == leftClean && value.RightClean == rightClean {
			currentState = value
			break
		}
	}
	return currentState
}
