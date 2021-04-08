package foundation

import "bpmn-validator/internals/models/core/infrastructure"


type BaseElement struct{
	id string 
	documentations []infrastructure.Documentation
	extensionDefinitions []ExtensionDefinition
	extensionValues []ExtensionAttributeValue
}