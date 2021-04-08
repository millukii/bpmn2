package bpmn

type NonInterruptingBoundaryMessageIntermediateEvent struct {
	id string
	name string
	attachToRef string
	cancelActivity bool //false
	messageEventDefinition MessageEventDefinition
}