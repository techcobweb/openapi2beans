package generator

type SchemaType struct {
	name        string
	path        string
	description string
	properties  []Property
}

type Property struct {
	name           string
	description    string
	typeName       string
	possibleValues []string
	resolvedType   SchemaType
	cardinality    Cardinality
}

type Cardinality struct {
	min int
	max int
}
