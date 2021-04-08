package common

type CorrelationProperty struct {
	name string
	correlationPropertyType string
	correlationPropertyRetrievalExpression []CorrelationPropertyRetrievalExpression
}