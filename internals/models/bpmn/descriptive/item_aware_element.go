package bpmn

type ItemAwareElement struct{
	ioSpecification string
	inputSet interface{}
	outputSet interface{}
	dataInput  interface{}
	dataOutput interface{}
}