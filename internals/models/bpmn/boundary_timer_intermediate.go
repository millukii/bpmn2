package bpmn

type BoundaryTimerIntermediateEvent struct {
	id string
	name string
	attachToRef string
	timerEventDefinition TimerEventDefinition
}