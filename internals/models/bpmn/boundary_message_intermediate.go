package bpmn

type BoundaryMessageIntermediate struct {
	id string
	name string
	attachToRef string
	messageEventDefinition MessageEventDefinition
}