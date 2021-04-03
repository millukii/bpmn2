package infrastructure

type Import struct{
	importType string
	location string
	namespace string
}

func NewImport(importType string, namespace string, location string) (*Import, error) {


	newImport := &Import{
		location:   location,
		namespace:  namespace,
	}

	newImport.SetImportType(importType)
	return newImport, nil
}

func(i *Import) SetImportType(importType  string) {
	switch importType {
	case "XML":
		i.importType= "https://www.w3.org/2001/XMLSchema"
	case "WSDL":
		i.importType= "http://www.w3.org/TR/wsdl20/"
	case "BPMN":
		i.importType= "BPMN"
	case "PDF":
		i.importType= "PDF"
	case "JPEG":
		i.importType= "JPEG"
	case "PNG":
		i.importType= "PNG"
	default://bpmn2.0
		i.importType= "http://www.omg.org/spec/BPMN/20100524/MODEL"
	}
}