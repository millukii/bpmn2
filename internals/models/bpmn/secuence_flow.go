package bpmn

type SecuenceFlow struct{
	id string
	name string
	sourceRef string
	targetRef string
}

type SecuenceFlowUnconditional struct{
	id string
	name string
	sourceRef string
	targetRef string
}