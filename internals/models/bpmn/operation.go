package bpmn

type Operation struct {
	id string
	name string
	inMessageRef Message
	outMessageRef Message
	errorsRefs []Error
}