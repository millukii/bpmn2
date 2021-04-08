package infrastructure

import (
	"bpmn-validator/internals/models/bpmn"
	"bpmn-validator/internals/models/core/foundation"
	"encoding/xml"
	"log"
)

//Infrastructure: Two elements that are used for both abstract syntax models and diagram models
type Infrastructure struct {
	XMLName        xml.Name `xml:"definitions"`
	Imports         []Import `xml:"import"`
	RootElements         []foundation.RootElement `xml:"rootElement"`
	BpmnDI []bpmn.BPMNDiagram `xml:"bpmndi:BPMNDiagram"`
	Name string `xml:"name,attr"`
	Location string `xml:"location,attr"`
	ImportType string `xml:"importType,attr"`
	TargetNameSpace string `xml:"targetNameSpace,attr"`
	ExpressionLanguage string  `xml:"expressionLanguage,attr"`
	TypeLanguage string   `xml:"typeLanguage,attr"`
	Exporter string  `xml:"exporter,attr"`
	ExporterVersion string  `xml:"exporterVersion,attr"`
	Relationships []foundation.Relationship `xml:"relationship"`
	Extension []foundation.Extension `xml:"extension"`
}

func NewInfraestructure() (*Infrastructure, error) {

	newInfra := &Infrastructure{}

	return newInfra, nil
}

func (i *Infrastructure) ToXML() (xml.CharData, error) {
	d, err := xml.Marshal(&i)
	if err != nil {
		log.Fatalf("xml.Marshal failed with '%s'\n", err)
		return nil, err
	}
	return xml.MarshalIndent(&d, "", "  ")
}
var xsd = `
<xsd:element name="definitions" type="tDefinitions"/>
<xsd:complexType name="tDefinitions">
<xsd:sequence>
<xsd:element ref="import" minOccurs="0" maxOccurs="unbounded"/>
<xsd:element ref="extension" minOccurs="0" maxOccurs="unbounded"/>
<xsd:element ref="rootElement" minOccurs="0" maxOccurs="unbounded"/>
<xsd:element ref="bpmndi:BPMNDiagram" minOccurs="0" maxOccurs="unbounded"/>
<xsd:element ref="relationship" minOccurs="0" maxOccurs="unbounded"/>
</xsd:sequence>
<xsd:attribute name="id" type="xsd:ID" use="optional"/>
<xsd:attribute name="targetNamespace" type="xsd:anyURI" use="required"/>
<xsd:attribute name="expressionLanguage" type="xsd:anyURI" use="optional" default="http://
www.w3.org/1999/XPath"/>
<xsd:attribute name="typeLanguage" type="xsd:anyURI" use="optional" default="http://www.w3.org/
2001/XMLSchema"/>
<xsd:anyAttribute name="exporter" type="xsd:ID"/>
<xsd:anyAttribute name="exporterVersion" type="xsd:ID"/>
<xsd:anyAttribute namespace="##other" processContents="lax"/>
</xsd:complexType>
<xsd:element name="import" type="tImport"/>
<xsd:complexType name="tImport">
<xsd:attribute name="namespace" type="xsd:anyURI" use="required"/>
<xsd:attribute name="location" type="xsd:string" use="required"/>
<xsd:attribute name="importType" type="xsd:anyURI" use="required"/>
</xsd:complexType>

`