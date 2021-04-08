package bpmn

type BoundaryErrorIntermediateEvent struct {
	id string
	name string
	attachToRef string
	errorEventDefinition ErrorEventDefinition
}