package bpmn

type NonInterruptingBoundaryConditionalIntermediateEvent struct  {
	id string
	name string
	cancelActivity bool //false
	conditionalEventDefinition ConditionalEventDefinition
}