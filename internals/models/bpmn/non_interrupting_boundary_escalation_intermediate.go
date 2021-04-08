package bpmn

type NonInterruptingBoundaryEscalationIntermediateEvent struct {
	id string
	name string
	attachedToRef string
	cancelActivity bool //false
	escalationEventDefinition EscalationEventDefinition
}