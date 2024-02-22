package generator

// TODO: Instead of type Object, Variable, we should be using *Object and *Variable.
// What we have here will pass by value rather than by reference...
// Possibly not overly important for performance in this application,
// but just so you know it's inefficient.

type Bean struct {
	// Would call it 'package' but that's a keyword in go!
	beanPackage string
	object      Object // TODO: Rename this to ObjectSchemaType? or something which isn't object.
}

type SchemaPart interface {
	// For arrays/lists, this is the type of each element.
	GetType() string

	GetName() string
	GetDescription() string

	// Is it an array or list of things ?
	IsMultipleCarinality() bool
}

type Variable struct {
	varDescription string

	// Would call it 'type' but that's a keyword in go!
	// For arrays/lists, this is the type of the variable/property/data member
	varTypeName string

	varName string

	// Is the variable multiple values ? ie: An array or list ?
	isMultipleCardinality bool

	// Do we want it to be set in the constructor ?
	isSetInConstructor bool
}

func (variable Variable) GetType() string {
	return variable.varTypeName
}

func (variable Variable) GetName() string {
	return variable.varName
}

func (variable Variable) GetDescription() string {
	return variable.varDescription
}

func (variable Variable) IsSetInConstructor() bool {
	return variable.isSetInConstructor
}

func (variable Variable) IsMultipleCarinality() bool {
	return variable.isMultipleCardinality
}

func (variable *Variable) SetType(varType string) {
	variable.varTypeName = varType
}

type Object struct {
	description string
	// Would call it 'type' but that's a keyword in go!
	varTypeName string
	varName     string
	variables   map[string]SchemaPart
}

func (obj Object) GetType() string {
	return obj.varTypeName
}

func (obj Object) GetName() string {
	return obj.varName
}

func (obj Object) GetDescription() string {
	return obj.description
}

func (obj Object) IsMultipleCarinality() bool {
	return false
}
