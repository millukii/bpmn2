package bpmn

/* a. Data Association is ABSTRACT: Data Input Association and Data Output Association will appear in
the XML serialization. These both have REQUIRED attributes [sourceRef and targetRef] which refer to
itemAwareElements. To be consistent with the metamodel, this will require the following additional
elements: ioSpecification, inputSet, outputSet, Data Input, Data Output. When a BPMN editor
draws a Data Association to an Activity or Event it should generate this supporting invisible substructure.
Otherwise, the metamodel would have to be changed to make sourceRef and targetRef optional or allow
reference to non-itemAwareElements, e.g., Activity and Event.
b. associationDirection not specified for Data Association */

type DataAssociation struct{
	id string
	name string
	sourceRef string //required
	targetRef string  //required
	associationDirection string
	assignment string
}

type DataInputAssociation struct{
	id string
	name string
	sourceRef string
	targetRef string
	associationDirection string
}
type DataOutputAssociation struct{
	id string
	name string
	sourceRef string
	targetRef string
	associationDirection string
}