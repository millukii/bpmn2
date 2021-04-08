package bpmn

type MultiInstanceLoopCharacteristics struct {
	id string
	isSequential bool
	loopDataInput interface{}
	inputDataItem interface{}
}