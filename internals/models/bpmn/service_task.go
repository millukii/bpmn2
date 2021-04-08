package bpmn

type ServiceTask struct{
	id string
	name string
	implementation string
	operationRef string
	ioSpecification string
	dataInputAssociations []DataInputAssociation
	dataOutputAssociations []DataOutputAssociation
	loopCharacteristics  string
	boundaryEventRefs string
}