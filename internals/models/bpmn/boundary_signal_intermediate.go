package bpmn

type BoundarySignalIntermediateEvent struct  {
	id string
	name string
	attachedToRef string
	signalEventDefinition SignalEventDefinition
}