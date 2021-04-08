package bpmn

type LaneSet struct {
	id   string
	lane []Lane
	//childLaneSet
	flowElementRef string
}
