package bpmn

type NonInterruptingBoundaryTimerIntermediateEvent struct {
	id string
	name string
	attachToRef string
	timerEventDefinition TimerEventDefinition
}