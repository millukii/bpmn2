package bpmn

type CatchingTimerIntermediateEvent struct {
	id string
	name string
	attachToRef string
	timerEventDefinition TimerEventDefinition
}