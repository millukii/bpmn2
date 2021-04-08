package bpmn

type MessageEndEvent struct{
	id string
	name string
	messageEventDefinition string
	dataInput DataInput
	dataInputAssociations []DataAssociation
}