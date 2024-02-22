package generator

type SchemaType struct {
	name        string
	path        string
	description string
}

type Property struct {
	name         string
	description  string
	typeName     string
	resolvedType SchemaType
	cardinality Cardinality
}

type Cardinality struct {
	min int
	max int
}

type PossibleValues struct {
	name string
	possibleValues []PossibleValue
}

type PossibleValue struct {
	value string
}