package infrastructure

import (
	"bpmn-validator/internals/models/bpmn"
	"bpmn-validator/internals/models/core/foundation"
)

type Definition struct {
	name string 
	targetNameSpace string
	expressionLanguage string
	typeLanguage string
	exporter string
	exporterVersion string
	rootElement []foundation.RootElement
	bpmnDiagrams []bpmn.BPMNDiagram
	imports []Import
	relationships []foundation.Relationship
	extension []foundation.Extension
 }

 func NewDefinition() (*Definition,error){

	newDefinition := &Definition{
		typeLanguage: "http://www.w3.org/2001/XMLSchema",
		exporter: "bpmn-validator",
		exporterVersion: "0.0.1",
	}
	return newDefinition, nil
 }