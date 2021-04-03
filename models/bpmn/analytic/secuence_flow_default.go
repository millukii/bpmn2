package analytic

type  SecuenceFlowDefaul struct{
	id string
	name string
	sourceRef  string
	targetRef string
	//Default is an attribute of a sourceRef (exclusive or inclusive) Gateway
	_default string
}