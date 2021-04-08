package bpmn

type  SecuenceFlowConditional struct{
	id string
	name string
	sourceRef  string
	targetRef string
	//ConditionExpression, allowed only for Sequence Flow out of Gateways, MAY b
	conditionExpresion string
}