package fsm

// EventProcessor defines OnExit, Action and OnEnter actions.
type EventProcessor interface {
	// OnExit Action handles exiting a state
	OnExit(fromState string, args []interface{})
	// Action is used to handle transitions
	Action(action string, fromState string, toState string, args []interface{}) error
	// OnExit Action handles entering a state
	OnEnter(toState string, args []interface{})
}

// DefaultDelegate is a default delegate.
// it splits processing of actions into three actions: OnExit, Action and OnEnter.
type DefaultDelegate struct {
	P EventProcessor
}

// HandleEvent implements Delegate interface and split HandleEvent into three actions.
func (dd *DefaultDelegate) HandleEvent(action string, fromState string, toState string, args []interface{}) {
	if fromState != toState {
		dd.P.OnExit(fromState, args)
	}

	if err := dd.P.Action(action, fromState, toState, args); err == nil {
		//如果执行action没有错误，才可以转换状态
		if fromState != toState {
			dd.P.OnEnter(toState, args)
		}
	}
}
