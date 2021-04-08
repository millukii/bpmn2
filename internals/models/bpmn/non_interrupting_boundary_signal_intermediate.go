package bpmn

type NonInterrumptingBoundarySignalIntermediateEvent struct  {
	id string
	name string
	attachedToRef string
	cancelActivity bool//false
	signalEventDefinition SignalEventDefinition
}