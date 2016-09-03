package task

// Task1State state for task1
type Task1State struct {
	position  string
	isClean   bool
	nextState *Task1State
}

var firstState = Task1State{"left", false, nil}

func init() {

}
