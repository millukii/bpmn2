package bpmn

type MessageStartEvent struct{
	id string
	name string
	messageEventDefinition string
	dataOutput DataOutput
	dataOutputAssociations []DataOutputAssociation
}
