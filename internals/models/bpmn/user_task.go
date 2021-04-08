package bpmn

type UserTask struct{
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
//todo checkout datatypes