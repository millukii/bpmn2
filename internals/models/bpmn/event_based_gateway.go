package bpmn

type EventBasedGateway struct {
	id string
	name string
	gatewayDirection GatewayDirection
	eventGatewayType EventGatewayType
}