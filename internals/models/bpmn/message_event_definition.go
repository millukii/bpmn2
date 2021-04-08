package bpmn

type MessageEventDefinition struct {
	id string
	messageRef Message
	operationRef Operation
}